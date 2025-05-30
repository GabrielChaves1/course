<template>
  <button
    class="flex items-center justify-center gap-2 px-6 py-3 rounded-full transition-colors"
    :class="[sizeClasses[size], variantClasses[variant]]"
    v-bind="$attrs"
  >
    <slot />
  </button>
</template>

<script setup lang="ts">
interface Props {
  variant?: "primary" | "secondary" | "danger";
  size?: "sm" | "md" | "lg";
}

withDefaults(defineProps<Props>(), {
  variant: "primary",
  size: "md",
});

const sizeClasses = {
  sm: "px-4 py-2 text-sm",
  md: "px-6 py-3 text-base",
  lg: "px-8 py-4 text-lg",
};

const variantClasses = {
  primary: "bg-indigo-600 text-white hover:bg-indigo-700 border-indigo-600",
  secondary: "bg-gray-600 text-white hover:bg-gray-700 border-gray-600",
  danger: "bg-red-600 text-white hover:bg-red-700 border-red-600",
};
</script>
