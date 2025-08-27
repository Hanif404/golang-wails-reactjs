import { useState } from 'react';
import { Container, Nav,Navbar, Offcanvas, NavDropdown} from 'react-bootstrap';
import { Link } from 'react-router-dom';

function Sidebar() {
    const [show, setShow] = useState(false);
    const handleClose = () => setShow(false);
    const handleShow = () => setShow(true);
    return (
        <>
            <Navbar expand="lg" className="bg-body-tertiary">
                <Navbar.Brand onClick={handleShow} style={{cursor: "pointer", paddingLeft:'20px'}}>My App</Navbar.Brand>
                <Container fluid>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="ms-auto">
                            <Nav.Link href="#login">{localStorage.getItem('user')}</Nav.Link>
                            <Nav.Link href="#signup">Log Out</Nav.Link>
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
            <Offcanvas show={show} onHide={handleClose} placement="start">
                <Offcanvas.Header closeButton>
                    <Offcanvas.Title>Sistem Sekolah</Offcanvas.Title>
                </Offcanvas.Header>
                <Offcanvas.Body>
                    <Nav className="flex-column">
                        <Nav.Link as={Link} to="/student" onClick={handleClose}>Student</Nav.Link>
                        <Nav.Link as={Link} to="/sync" onClick={handleClose}>Sync</Nav.Link>
                    </Nav>
                </Offcanvas.Body>
            </Offcanvas>
        </>
    )
}

export default Sidebar