<template>
  <div class="flex items-center justify-center gap-4">
    <div v-for="(group, groupIndex) in code" :key="groupIndex" class="flex gap-2">
      <input v-for="(char, charIndex) in group" :key="`code-${groupIndex}-${charIndex}`"
        :ref="(el) => setInputRef(el as HTMLInputElement, groupIndex, charIndex)" type="text" inputmode="numeric"
        pattern="[0-9]*" maxlength="1" autocomplete="one-time-code"
        class="w-12 h-12 text-center text-xl font-semibold border-2 border-gray-300 rounded-lg bg-white transition-all duration-200 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500 hover:border-gray-400 disabled:bg-gray-100 disabled:cursor-not-allowed"
        :class="{
          'border-red-500 focus:ring-red-500 focus:border-red-500': hasError && char,
          'border-green-500': isComplete && char
        }" :value="code[groupIndex][charIndex]" :disabled="disabled" @input="onInput($event, groupIndex, charIndex)"
        @keydown="onKeyDown($event, groupIndex, charIndex)" @paste="onPaste($event, groupIndex, charIndex)"
        @focus="onFocus(groupIndex, charIndex)" />
    </div>
    <div v-if="groupIndex < code.length - 1" v-for="groupIndex in code.length - 1" :key="`separator-${groupIndex}`"
      class="w-3 h-0.5 bg-gray-300 rounded"></div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, nextTick, computed } from 'vue'

interface Props {
  modelValue: string
  length?: number
  disabled?: boolean
  hasError?: boolean
  allowAlphanumeric?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  length: 6,
  disabled: false,
  hasError: false,
  allowAlphanumeric: false
})

const emit = defineEmits<{
  'update:modelValue': [value: string]
  'complete': [value: string]
  'change': [value: string]
}>()

const groupCount = computed(() => Math.ceil(props.length / 3))
const charsPerGroup = computed(() => Math.ceil(props.length / groupCount.value))

const code = ref<string[][]>([])
const inputRefs = ref<HTMLInputElement[][]>([])

const isComplete = computed(() => {
  return code.value.flat().every(char => char.trim() !== '') &&
    code.value.flat().length === props.length
})

function initializeCode() {
  const groups: string[][] = []
  const refs: HTMLInputElement[][] = []

  for (let i = 0; i < groupCount.value; i++) {
    const groupSize = i === groupCount.value - 1
      ? props.length - (i * charsPerGroup.value)
      : charsPerGroup.value

    groups.push(new Array(groupSize).fill(''))
    refs.push([])
  }

  code.value = groups
  inputRefs.value = refs
}

initializeCode()

function setInputRef(el: HTMLInputElement | null, group: number, index: number) {
  if (el) {
    if (!inputRefs.value[group]) inputRefs.value[group] = []
    inputRefs.value[group][index] = el
  }
}

watch(() => props.modelValue, (val) => {
  const cleanValue = (val || '').replace(/\s/g, '')
  const padded = cleanValue.padEnd(props.length, '')

  let charIndex = 0
  for (let groupIndex = 0; groupIndex < code.value.length; groupIndex++) {
    for (let i = 0; i < code.value[groupIndex].length; i++) {
      code.value[groupIndex][i] = padded[charIndex] || ''
      charIndex++
    }
  }
}, { immediate: true })

watch(code, () => {
  const fullCode = code.value.flat().join('')
  emit('update:modelValue', fullCode)
  emit('change', fullCode)

  if (isComplete.value) {
    emit('complete', fullCode)
  }
}, { deep: true })

function validateInput(value: string): boolean {
  if (props.allowAlphanumeric) {
    return /^[a-zA-Z0-9]$/.test(value)
  }
  return /^[0-9]$/.test(value)
}

function onInput(event: Event, group: number, index: number) {
  const target = event.target as HTMLInputElement
  const value = target.value.slice(-1)

  if (value && !validateInput(value)) {
    target.value = code.value[group][index]
    return
  }

  code.value[group][index] = value

  if (value) {
    moveToNext(group, index)
  }
}

function onKeyDown(event: KeyboardEvent, group: number, index: number) {
  if (event.key === 'Backspace') {
    event.preventDefault()

    if (code.value[group][index]) {
      code.value[group][index] = ''
    } else {
      moveToPrevious(group, index)
    }
  } else if (event.key === 'Delete') {
    event.preventDefault()
    code.value[group][index] = ''
  } else if (event.key === 'ArrowLeft') {
    event.preventDefault()
    moveToPrevious(group, index)
  } else if (event.key === 'ArrowRight') {
    event.preventDefault()
    moveToNext(group, index)
  } else if (event.key === 'Home') {
    event.preventDefault()
    focus(0, 0)
  } else if (event.key === 'End') {
    event.preventDefault()
    const lastGroup = code.value.length - 1
    const lastIndex = code.value[lastGroup].length - 1
    focus(lastGroup, lastIndex)
  }
}

function onPaste(event: ClipboardEvent) {
  event.preventDefault()

  const pastedData = event.clipboardData?.getData('text') || ''
  const cleanData = pastedData.replace(/\D/g, '').slice(0, props.length)

  if (cleanData) {
    code.value.forEach(group => group.fill(''))

    let charIndex = 0
    for (let groupIndex = 0; groupIndex < code.value.length && charIndex < cleanData.length; groupIndex++) {
      for (let i = 0; i < code.value[groupIndex].length && charIndex < cleanData.length; i++) {
        if (validateInput(cleanData[charIndex])) {
          code.value[groupIndex][i] = cleanData[charIndex]
        }
        charIndex++
      }
    }

    const nextEmpty = findNextEmptyField()
    if (nextEmpty) {
      focus(nextEmpty.group, nextEmpty.index)
    }
  }
}

function onFocus(group: number, index: number) {
  nextTick(() => {
    const input = inputRefs.value[group]?.[index]
    if (input) {
      input.select()
    }
  })
}

function moveToNext(group: number, index: number) {
  if (index < code.value[group].length - 1) {
    focus(group, index + 1)
  } else if (group < code.value.length - 1) {
    focus(group + 1, 0)
  }
}

function moveToPrevious(group: number, index: number) {
  if (index > 0) {
    focus(group, index - 1)
  } else if (group > 0) {
    const prevGroupLastIndex = code.value[group - 1].length - 1
    focus(group - 1, prevGroupLastIndex)
  }
}

function focus(group: number, index: number) {
  nextTick(() => {
    const input = inputRefs.value[group]?.[index]
    if (input) {
      input.focus()
    }
  })
}

function findNextEmptyField(): { group: number; index: number } | null {
  for (let groupIndex = 0; groupIndex < code.value.length; groupIndex++) {
    for (let charIndex = 0; charIndex < code.value[groupIndex].length; charIndex++) {
      if (!code.value[groupIndex][charIndex]) {
        return { group: groupIndex, index: charIndex }
      }
    }
  }
  return null
}

defineExpose({
  focus: () => focus(0, 0),
  clear: () => {
    code.value.forEach(group => group.fill(''))
    focus(0, 0)
  },
  setValue: (value: string) => {
    emit('update:modelValue', value)
  }
})
</script>

<style scoped>
input::-webkit-outer-spin-button,
input::-webkit-inner-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}

input {
  transition: all 0.2s ease-in-out;
}

input.border-red-500 {
  animation: shake 0.3s ease-in-out;
}

@keyframes shake {

  0%,
  100% {
    transform: translateX(0);
  }

  25% {
    transform: translateX(-2px);
  }

  75% {
    transform: translateX(2px);
  }
}
</style>
