import {createRoot} from 'react-dom/client'
import './style.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import Login from './pages/login';
import react from 'react';

document.addEventListener('contextmenu', function(event) {
    event.preventDefault();
});

const container = document.getElementById('loginView')
const root = createRoot(container!)

root.render(
    <react.StrictMode>
        <Login />
    </react.StrictMode>
)
