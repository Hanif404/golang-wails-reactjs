import { useState } from 'react';
import { Container, Row, Col, Form, Button, Alert} from 'react-bootstrap';
import { ActionLogin } from "../../wailsjs/go/auth/App";
import { auth } from "../../wailsjs/go/models";

function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [isSubmit, setIsSubmit] = useState(false);
    const [alert, setAlert] = useState(false);
    const [message, setMessage] = useState('');
    const updateEmail = (e: any) => setEmail(e.target.value);
    const updatePassword = (e: any) => setPassword(e.target.value);
    
    if(localStorage.getItem('user') && localStorage.getItem('user') != ""){
        window.location.href = "./main.html";
    }

    const submitLogin = (event: any) => {
        event.preventDefault();
        setIsSubmit(true);

        let login = new auth.Login();
        login.email = email;
        login.password = password;
        ActionLogin(login).then((result) => {
            setIsSubmit(false);

            const jsonResult = JSON.parse(result);
            if (jsonResult.status) {
                localStorage.setItem('user', email);
                window.location.href = "./main.html"; 
            }
            
            if (!jsonResult.status) {
                setAlert(true)
                setMessage(jsonResult.message)
            }
        });
    }

    return (
        <Container fluid>
            <Row>
                <Col sm={3} md={4}></Col>
                <Col style={{marginTop: '20%'}}>
                    {alert && 
                        <Alert key='danger' variant='danger'>
                            {message}
                        </Alert>
                    }
                    <Form onSubmit={submitLogin}>
                        <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                            <Form.Label>Email</Form.Label>
                            <Form.Control autoComplete="off" onChange={updateEmail} name="input" type="email" placeholder="enter email" />
                        </Form.Group>
                        <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
                            <Form.Label>Password</Form.Label>
                            <Form.Control autoComplete="off" onChange={updatePassword} name="input" type="password" placeholder="enter password"/>
                        </Form.Group>
                        <Button variant="primary" type="submit" className="w-100" disabled={isSubmit}>
                            Log In
                        </Button>
                    </Form>
                </Col>
                <Col sm={3} md={4}></Col>
            </Row>
        </Container>
    )
}

export default Login