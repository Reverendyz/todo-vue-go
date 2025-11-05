// useTodos.ts
import { ref, onMounted } from 'vue'
import { getTasks, addTask, toggleTask, type Task } from '@/api/tasks'

export function useTodos() {
  const todos = ref<Task[] | null>(null)
  const loading = ref(false)

  async function load() {
    loading.value = true
    try {
      const data = await getTasks()
      todos.value = Array.isArray(data) ? data : []
    } finally {
      loading.value = false
    }
  }

  async function create(taskText: string) {
    if (!todos.value) {
      todos.value = []
    }

    const tempId = crypto.randomUUID()
    const tempTask: Task = { _id: tempId, task: taskText, done: false }
    todos.value.push(tempTask)

    try {
      const id = await addTask(taskText)
      const task = todos.value.find(t => t._id === tempId)
      if (task) {
        task._id = id
      }

      return id
    } catch (err) {
      todos.value = todos.value.filter(t => t._id !== tempId)
      throw err
    }
  }

  async function toggle(id: string) {
    if (!todos.value) {
      todos.value = []
    }
    const t = todos.value.find(x => x._id === id)
    if (!t) throw new Error('task not found')
    const old = t.done
    t.done = !old
    try {
      await toggleTask(id)
    } catch (err) {
      t.done = old
      throw err
    }
  }
  onMounted(load)
  return { todos, loading, load, create, toggle }
}
