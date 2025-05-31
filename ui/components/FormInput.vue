<template>
  <div class="w-full">
    <label :for="id" class="block text-sm font-medium text-gray-600">{{ label }}</label>

    <div class="relative">
      <input :id="id" :type="inputType" :value="modelValue" @input="$emit('update:modelValue', $event.target.value)"
        :class="[
          'mt-1 block w-full pr-10 px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500',
          error && 'border-red-500'
        ]" :placeholder="placeholder" />

      <button v-if="type === 'password'" type="button" class="absolute right-2 top-2.5 text-gray-400"
        @click="toggleVisibility" tabindex="-1">
        <Icon v-if="inputType === 'password'" name="material-symbols:visibility-rounded" class="size-5 text-gray-400" />
        <Icon v-else name="material-symbols:visibility-off-rounded" class="size-5 text-gray-400" />
      </button>
    </div>

    <p v-if="error" class="text-red-600 text-sm mt-1">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface Props {
  id: string;
  label: string;
  type?: string;
  modelValue: string;
  placeholder?: string;
  error?: string;
}

const props = defineProps<Props>()
defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>()

const isVisible = ref(false)
const inputType = computed(() => {
  return props.type === 'password' && isVisible.value ? 'text' : props.type || 'text'
})

function toggleVisibility() {
  isVisible.value = !isVisible.value
}
</script>
