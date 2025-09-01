import { Container, Form, Row, Col, Button, Alert } from "react-bootstrap";
import { useState } from "react";
import { ActionForm, Student } from "../../wailsjs/go/student/App";
import { student as studentModel } from "../../wailsjs/go/models";
import { useNavigate, useParams } from "react-router-dom";

function StudentForm() {
    const navigate = useNavigate();
    const { id } = useParams();

    const [validated, setValidated] = useState(false);
    const [alert, setAlert] = useState(false);
    const [message, setMessage] = useState('');
    const [nis, setNIS] = useState('');
    const [name, setName] = useState('');
    const [gender, setGender] = useState('');
    const [tempatLahir, setTempatLahir] = useState('');
    const [tanggalLahir, setTanggalLahir] = useState('');
    const [kelas, setKelas] = useState('');
    const [angkatan, setAngkatan] = useState('');
    const [alamat, setAlamat] = useState('');
    const [namaAyah, setNamaAyah] = useState('');
    const [namaIbu, setNamaIbu] = useState('');
    const [phone, setPhone] = useState('');
    const [email, setEmail] = useState('');
    const [status, setStatus] = useState('');

    const updateNis = (e: any) => setNIS(e.target.value);
    const updateName = (e: any) => setName(e.target.value);
    const updateGender = (e: any) => setGender(e.target.value);
    const updateTempat = (e: any) => setTempatLahir(e.target.value);
    const updateTanggal = (e: any) => setTanggalLahir(e.target.value);
    const updateKelas = (e: any) => setKelas(e.target.value);
    const updateAngkatan = (e: any) => setAngkatan(e.target.value);
    const updateAlamat = (e: any) => setAlamat(e.target.value);
    const updateNameAyah = (e: any) => setNamaAyah(e.target.value);
    const updateNameIbu = (e: any) => setNamaIbu(e.target.value);
    const updatePhone = (e: any) => setPhone(e.target.value);
    const updateEmail = (e: any) => setEmail(e.target.value);
    const updateStatus = (e: any) => setStatus(e.target.value);

    if(id){
        Student(Number(id)).then(result => {
            const jsonResult = JSON.parse(result);
            if(jsonResult.status){
                const data = jsonResult.data;
                setNIS(data.nis);
                setName(data.name);
                setGender(data.gender);
                setTempatLahir(data.tempat_lahir);
                setTanggalLahir(data.tanggal_lahir);
                setKelas(data.kelas);
                setAngkatan(data.angkatan);
                setAlamat(data.alamat);
                setNamaAyah(data.nama_ayah);
                setNamaIbu(data.nama_ibu);
                setPhone(data.phone);
                setEmail(data.email);
                setStatus(data.status);
            }
        })
    }

    const handleSubmit = (event: any) => {
        const form = event.currentTarget;
        if (form.checkValidity() === false) {
            event.preventDefault();
            event.stopPropagation();
        }
        setValidated(true);

        let student = new studentModel.Student();
        student.nis = nis;
        student.name = name;
        student.gender = gender;
        student.tempat_lahir = tempatLahir;
        student.tanggal_lahir = tanggalLahir;
        student.kelas = kelas;
        student.angkatan = angkatan;
        student.alamat = alamat;
        student.nama_ayah = namaAyah;
        student.nama_ibu = namaIbu;
        student.phone = phone;
        student.email = email;
        student.status = status;
        student.created_by = localStorage.getItem('user') || "";
        ActionForm(student).then((result) => {
            const jsonResult = JSON.parse(result);
            if (jsonResult.status) {
                navigate('/student');
            }else if (typeof jsonResult.message === 'string'){
                setAlert(true);
                setMessage(jsonResult.message);
            }
        });
    };

    const handleCancel = () =>{
        navigate('/student')
    }
    return (
        <Container fluid className="p-5">
            <h3>Form siswa</h3>
            {alert && 
                <Alert key='danger' variant='danger'>
                    {message}
                </Alert>
            }
            <Form noValidate validated={validated} onSubmit={handleSubmit}>
                <Form.Group as={Row} className="mb-3" controlId="form1">
                    <Form.Label column sm="2">NIS</Form.Label>
                    <Col sm="3">
                        <Form.Control type="text" placeholder="Nomor Induk Siswa" value={nis} onChange={updateNis} required/>
                        <Form.Control.Feedback type="invalid">
                            NIS wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form2">
                    <Form.Label column sm="2">Nama Lengkap</Form.Label>
                    <Col sm="10">
                        <Form.Control type="text" placeholder="Nama Lengkap" value={name} onChange={updateName} required/>
                        <Form.Control.Feedback type="invalid">
                            Nama Lengkap wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form3">
                    <Form.Label column sm="2">Jenis Kelamin</Form.Label>
                    <Col sm="2">
                        <Form.Control type="text" placeholder="L\P" value={gender} onChange={updateGender} required/>
                        <Form.Control.Feedback type="invalid">
                            Jenis Kelamin wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form4">
                    <Form.Label column sm="2">Tempat Lahir</Form.Label>
                    <Col sm="10">
                        <Form.Control type="text" placeholder="Tempat Lahir" value={tempatLahir} onChange={updateTempat} required/>
                        <Form.Control.Feedback type="invalid">
                            Tempat Lahir wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form5">
                    <Form.Label column sm="2">Tanggal Lahir</Form.Label>
                    <Col sm="3">
                        <Form.Control type="date" placeholder="Tanggal Lahir" value={tanggalLahir} onChange={updateTanggal} required/>
                        <Form.Control.Feedback type="invalid">
                            Tanggal Lahir wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form6">
                    <Form.Label column sm="2">Kelas</Form.Label>
                    <Col sm="3">
                        <Form.Control type="text" placeholder="Kelas" value={kelas} onChange={updateKelas} required/>
                        <Form.Control.Feedback type="invalid">
                            Kelas wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form7">
                    <Form.Label column sm="2">Angkatan</Form.Label>
                    <Col sm="3">
                        <Form.Control type="text" placeholder="Angkatan" value={angkatan} onChange={updateAngkatan} required/>
                        <Form.Control.Feedback type="invalid">
                            Angkatan wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form8">
                    <Form.Label column sm="2">Alamat</Form.Label>
                    <Col sm="10">
                        <Form.Control type="text" placeholder="Alamat" value={alamat} onChange={updateAlamat} required/>
                        <Form.Control.Feedback type="invalid">
                            Alamat wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form8">
                    <Form.Label column sm="2">Nama Ayah</Form.Label>
                    <Col sm="10">
                        <Form.Control type="text" placeholder="Nama Ayah" value={namaAyah} onChange={updateNameAyah} required/>
                        <Form.Control.Feedback type="invalid">
                            Nama Ayah wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form8">
                    <Form.Label column sm="2">Nama Ibu</Form.Label>
                    <Col sm="10">
                        <Form.Control type="text" placeholder="Nama Ibu" value={namaIbu} onChange={updateNameIbu} required/>
                        <Form.Control.Feedback type="invalid">
                            Nama Ibu wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form8">
                    <Form.Label column sm="2">No. Telp</Form.Label>
                    <Col sm="10">
                        <Form.Control type="text" placeholder="No. Telp" value={phone} onChange={updatePhone} required/>
                        <Form.Control.Feedback type="invalid">
                            No. Telp wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form8">
                    <Form.Label column sm="2">Email</Form.Label>
                    <Col sm="10">
                        <Form.Control type="email" placeholder="Email" value={email} onChange={updateEmail} required/>
                        <Form.Control.Feedback type="invalid">
                            Email wajib diisi.
                        </Form.Control.Feedback>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="form8">
                    <Form.Label column sm="2">Status</Form.Label>
                    <Col sm="10">
                        <Form.Select aria-label="Status" onChange={updateStatus} value={status} required >
                            <option>Pilih Status</option>
                            <option value="aktif">Aktif</option>
                            <option value="non-aktif">Tidak Aktif</option>
                        </Form.Select>
                    </Col>
                </Form.Group>

                <Form.Group as={Row} className="mb-3" controlId="formNamaLengkap">
                    <Form.Label column sm="2">&nbsp;</Form.Label>
                    <Col sm="10">
                        <Button type="submit">Simpan</Button> &nbsp; 
                        <Button variant="danger" type="button" onClick={handleCancel}>Batal</Button>
                    </Col>
                </Form.Group>
            </Form>
        </Container>
    );
}

export default StudentForm;