import { useState } from 'react';
import { Container, Nav,Navbar, Offcanvas} from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { ActionLogout } from "../../../wailsjs/go/auth/App";

function Sidebar() {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);

    const submitLogout = (event: any) => {
        event.preventDefault();

        ActionLogout(localStorage.getItem('user') || '').then(() => {
            localStorage.removeItem('user');
            window.location.href = "./index.html";
        });
    }

    return (
        <>
            <Navbar expand="lg" className="bg-body-tertiary">
                <Navbar.Brand onClick={handleShow} style={{cursor: "pointer", paddingLeft:'20px'}}>Menu</Navbar.Brand>
                <Container fluid>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="ms-auto">
                            <Nav.Link href="#login">{localStorage.getItem('user')}</Nav.Link>
                            <Nav.Link href="#logout" onClick={submitLogout}>Log Out</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
            <Offcanvas show={show} onHide={handleClose} placement="start">
                <Offcanvas.Header closeButton>
                    <Offcanvas.Title>Sistem Manajemen Sekolah</Offcanvas.Title>
                </Offcanvas.Header>
                <Offcanvas.Body>
                    <Nav className="flex-column">
                        <Nav.Link as={Link} to="/" onClick={handleClose}>Home</Nav.Link>
                        <Nav.Link as={Link} to="/student" onClick={handleClose}>Data Siswa</Nav.Link>
                        <Nav.Link as={Link} to="/teacher" onClick={handleClose}>Data Guru & Staff</Nav.Link>
                        <Nav.Link as={Link} to="/absensi" onClick={handleClose}>Absensi</Nav.Link>
                        <Nav.Link as={Link} to="/kurikulum" onClick={handleClose}>Kurikulum</Nav.Link>
                        <Nav.Link as={Link} to="/schedule" onClick={handleClose}>Jadwal & Pembelajaran</Nav.Link>
                        <Nav.Link as={Link} to="/teacher" onClick={handleClose}>Penilaian</Nav.Link>
                        <Nav.Link as={Link} to="/student/payment" onClick={handleClose}>Tagihan Siswa</Nav.Link>
                        <Nav.Link as={Link} to="/teacher/payment" onClick={handleClose}>Gaji Guru & Staff</Nav.Link>
                        <Nav.Link as={Link} to="/inventory" onClick={handleClose}>Inventaris Sekolah</Nav.Link>
                        <Nav.Link as={Link} to="/tahfidz" onClick={handleClose}>Tahfidz Tahsin</Nav.Link>
                        <Nav.Link as={Link} to="/konseling" onClick={handleClose}>Bimbingan Konseling</Nav.Link>
                        <Nav.Link as={Link} to="/sync/upload" onClick={handleClose}>Sync Data</Nav.Link>
                    </Nav>
                </Offcanvas.Body>
            </Offcanvas>
        </>
    )
}

export default Sidebar