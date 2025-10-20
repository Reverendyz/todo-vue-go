<template>
  <label class="todo-item">
    <v-checkbox :class="{ checked: todo.done }" :model-value="todo.done" :label="todo.task" @update:model-value="emitToggle" />
  </label>
</template>

<script setup lang="ts">
import type { Task } from '@/api/tasks'

const props = defineProps<{ todo: Task }>()
const emit = defineEmits<{
  (e: 'toggle', id: string): void
}>()

function emitToggle() {
  emit('toggle', props.todo._id)
}
</script>

<style scoped>
.checked :deep(.v-label) {
  text-decoration: line-through;
}
</style>
