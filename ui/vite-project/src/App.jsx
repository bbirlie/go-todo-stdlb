import { useState } from 'react'
import {
  QueryClient,
  QueryClientProvider
} from '@tanstack/react-query'
import Todos from './components/todos'
import PostForm from './components/postform'

const queryClient = new QueryClient()


function App() {
  const [count, setCount] = useState(0)

  return (
    <QueryClientProvider client={queryClient}>
      <>
        <Todos></Todos>
        <PostForm></PostForm>
      </>
    </QueryClientProvider>
  )
}

export default App
