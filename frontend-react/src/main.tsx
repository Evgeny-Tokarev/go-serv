import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './style/index.css'
import store from './app/store'
import { Provider } from 'react-redux'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
      <Provider store={store} children={undefined}/>
    <App />
  </React.StrictMode>,
)
