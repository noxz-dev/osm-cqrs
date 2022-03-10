<script setup lang="ts">
const emit = defineEmits(['update:modelValue']);

const props = defineProps({
  modelValue: [String, Number]
});

const onChanged = (payload: Event) => {
  const event = payload as InputEvent;
  emit('update:modelValue', (event.target as HTMLInputElement).value);
};
</script>

<template>
  <input
    type="text"
    class="focus:outline-none border bg-gray-100 border-slate-300 rounded-md px-3 py-2"
    :value="modelValue"
    @input="onChanged"
  />
</template>
