<template>
  <component :is="as" :class="[variantClass, $attrs.class]" v-bind="$attrs">
    <slot />
  </component>
</template>

<script setup lang="ts">
interface Props {
  as?: string;
  variant?: "default" | "secondary";
}

const props = withDefaults(defineProps<Props>(), {
  as: "p",
  variant: "default",
  size: "base",
});

const variantClasses = {
  default: "text-gray-800",
  secondary: "text-gray-600",
};

const variantClass = computed(
  () => variantClasses[props.variant] || variantClasses.default,
);
</script>
