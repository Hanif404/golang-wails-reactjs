import react from 'react';
import {createRoot} from 'react-dom/client'
import './style.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import { Routes, Route, HashRouter } from 'react-router-dom';
import Sidebar from './pages/components/sidebar';
import Home from './pages/home';
import Sync from './pages/sync-upload';
import Student from './pages/student';
import StudentForm from './pages/student-form';
import Login from './pages/login';

// document.addEventListener('contextmenu', function(event) {
//     event.preventDefault();
// });

const container = document.getElementById('root')
const root = createRoot(container!)
const CommingSoon = () => {
    return (
        <div className="p-3">
            <h1>Comming Soon</h1>
        </div>
    )
}

if(localStorage.getItem('user') && localStorage.getItem('user') != ""){
    root.render(
        <>
        <HashRouter>
            <Sidebar />
            <Routes>
                <Route path="*" element={<Home />} />
                <Route path="/student" element={<Student />} />
                <Route path="/student/form" element={<StudentForm />} />
                <Route path="/student/form/:id" element={<StudentForm />} />
                <Route path="/teacher" element={<CommingSoon />} />
                <Route path="/absensi" element={<CommingSoon />} />
                <Route path="/kurikulum" element={<CommingSoon />} />
                <Route path="/schedule" element={<CommingSoon />} />
                <Route path="/student/payment" element={<CommingSoon />} />
                <Route path="/teacher/payment" element={<CommingSoon />} />
                <Route path="/inventory" element={<CommingSoon />} />
                <Route path="/tahfidz" element={<CommingSoon />} />
                <Route path="/konseling" element={<CommingSoon />} />
                <Route path="/sync/upload" element={<Sync />} />
            </Routes>
        </HashRouter>
        </>
    )
}else{
    root.render(
        <react.StrictMode>
            <Login />
        </react.StrictMode>
    )
}
