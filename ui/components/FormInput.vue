<template>
  <div class="mb-4 w-full">
    <label :for="id" class="block text-sm font-medium text-gray-600">{{ label }}</label>
    <input
      :id="id"
      :type="type"
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      :class="[
        'mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500',
        error && 'border-red-500'
      ]"
      :placeholder="placeholder"
    />
    <p v-if="error" class="text-red-600 text-sm mt-1">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
interface Props {
  id: string;
  label: string;
  type?: string;
  modelValue: string;
  placeholder?: string;
  error?: string;
}

defineProps<Props>();

defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();
</script>

