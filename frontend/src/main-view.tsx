import {createRoot} from 'react-dom/client'
import './style.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Sidebar from './pages/components/sidebar';
import Home from './pages/home';
import Sync from './pages/sync';
import Student from './pages/student';

// document.addEventListener('contextmenu', function(event) {
//     event.preventDefault();
// });

const container = document.getElementById('root')
const root = createRoot(container!)

root.render(
    <>
    <BrowserRouter>
        <Sidebar />
        <Routes>
            <Route path="*" element={<Home />} />
            <Route path="/student" element={<Student />} />
            <Route path="/sync" element={<Sync />} />
        </Routes>
    </BrowserRouter>
    </>
)
