import { useState } from 'react'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <h1>Hello Worlds {count}</h1>

      <button onClick={() => setCount(pre => pre + 1)}>click count +</button> <br />
      <button disabled={count === 0} onClick={() => setCount(pre => pre - 1)}>click count -</button>
    </>
  )
}

export default App
