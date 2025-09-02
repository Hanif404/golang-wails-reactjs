import { Container, Row, Col, Button } from "react-bootstrap"
import DataTable from 'datatables.net-react';
import DT from 'datatables.net-bs5';
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { Students, ActionDelete } from "../../wailsjs/go/student/App";

DataTable.use(DT);

function Student(){
    const navigate = useNavigate();
    const [tableData, setTableData] = useState([]);

    useEffect(() => {
        Students().then(result => {
            const jsonResult = JSON.parse(result);
            if(jsonResult.status){
                setTableData(jsonResult.data);
            }
        })
    }, []);

    const addForm = () => {
        navigate('/student/form');
    }

    const editForm = (event: any) => {
        const id = event.target.getAttribute('data-id');
        navigate(`/student/form/${id}`);
    }

    const deleteForm = (event: any) => {
        const id = event.target.getAttribute('data-id');
        ActionDelete(Number(id)).then((result: any) => {
            const jsonResult = JSON.parse(result);
            if(jsonResult.status){
                window.location.reload();
            }
        })
    }

    const columns = [
        { data: 'nis' },
        { data: 'name' },
        { data: 'gender' },
        { data: 'kelas' },
        { data: 'angkatan' },
        { data: 'nama_ayah' },
        { data: 'nama_ibu' },
        { data: 'phone' },
        { data: 'status' },
        { data: null }
    ];
    return(
        <Container fluid className="mt-3">
            <Row>
                <Col><h3>Data Siswa</h3></Col>
                <Col className="d-flex justify-content-end"><Button variant="primary" onClick={addForm}>Add</Button></Col>
            </Row>
            <DataTable data={tableData} columns={columns} 
                slots={{
                    9: (data: any, row: any) => (
                        <>
                            <Button onClick={editForm} data-id={data.id} size="sm">
                                Ubah
                            </Button> &nbsp;
                            <Button variant="danger" onClick={deleteForm} data-id={data.id} size="sm">
                                Hapus
                            </Button>
                        </>
                    )
                }}
            className="table table-sm table-bordered">
                <thead>
                    <tr>
                        <th>NIS</th>
                        <th>Nama Lengkap</th>
                        <th>Jenis Kelamin</th>
                        <th>Kelas</th>
                        <th>Angkatan</th>
                        <th>Nama Ayah</th>
                        <th>Nama Ibu</th>
                        <th>No. Telp</th>
                        <th>Status</th>
                        <th>&nbsp;</th>
                    </tr>
                </thead>
            </DataTable>
        </Container>
    )
}
export default Student