<script setup lang="ts">
import { ref } from 'vue'
import TodoItem from '@/components/TodoItem.vue'
import { useTodos } from '@/composables/useTodos'

const { todos, create, toggle } = useTodos()

const newTask = ref('')

async function addItem() {
  const text = newTask.value.trim()
  if (!text) return

  try {
    await create(text)
    newTask.value = ''
  } catch (err) {
    console.error('Error while creating task:', err)
  }
}
async function onToggle(id: string) {
  try {
    await toggle(id)
  } catch (err) {
    console.error('Error on toggle task', err)
  }
}
</script>

<template>
  <v-container class="d-flex overflow-auto h-50">
    <v-col>
      <v-row v-for="todo in todos" :key="todo._id">
        <todo-item :todo="todo" @toggle="onToggle" />
      </v-row>
    </v-col>
  </v-container>
  <v-divider></v-divider>
  <v-container class="d-flex ga-1">
    <v-row>
      <v-col>
        <v-text-field
          clearable
          label="Insert your todo"
          variant="underlined"
          class="mx-auto"
          max-width="300"
          v-model.lazy.trim="newTask"
        />
      </v-col>
      <v-col>
        <v-btn prepend-icon="mdi-plus" @click.stop="addItem"> Add todo</v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>
