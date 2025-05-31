<template>
  <div class="w-full">
    <label :for="id" class="block text-sm font-medium text-gray-600">{{ label }}</label>
    <div class="flex gap-2">
      <select v-model="selectedCode" @change="emitCombined" :class="[
        'mt-1 block w-1/3 pr-10 px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500',
        error ? 'border-red-500' : 'border-gray-300'
      ]">
        <option v-for="code in countryCodes" :key="code.value" :value="code.value">
          {{ code.label }}
        </option>
      </select>

      <input :id="id" type="tel" v-model="localNumber" @input="emitCombined" :class="[
        'mt-1 block w-2/3 pr-10 px-3 py-2 border rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500',
        error ? 'border-red-500' : 'border-gray-300'
      ]" :placeholder="placeholder" />
    </div>
    <p v-if="error" class="text-red-600 text-sm mt-1">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, defineProps, defineEmits } from 'vue'

const props = defineProps<{
  id: string
  label: string
  modelValue: string
  placeholder?: string
  error?: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const countryCodes = [
  { label: '+55', value: '+55' },
  { label: '+1', value: '+1' },
  { label: '+44', value: '+44' },
  { label: '+33', value: '+33' },
]

const selectedCode = ref('+55')
const localNumber = ref('')

watch(
  () => props.modelValue,
  (newVal) => {
    const matchedCode = countryCodes.find((code) => newVal.startsWith(code.value))
    selectedCode.value = matchedCode?.value || '+55'
    localNumber.value = newVal.replace(selectedCode.value, '')
  },
  { immediate: true }
)

function emitCombined() {
  emit('update:modelValue', selectedCode.value + localNumber.value)
}
</script>
