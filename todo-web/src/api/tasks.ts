import { api } from './api'

export type Task = {
  _id: string
  task: string
  done: boolean
}

export async function getTasks(): Promise<Task[]> {
  const res = await api.get('/')
  return res.data.Found
}

export async function addTask(task: string): Promise<string> {
  const res = await api.post('/', { task })
  return res.data.message
}

export async function toggleTask(id: string) {
  return api.patch(`/${id}`)
}
