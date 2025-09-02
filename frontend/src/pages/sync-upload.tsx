import { Container, Row, Col, Button } from "react-bootstrap"
import DataTable from 'datatables.net-react';
import DT from 'datatables.net-bs5';
import { useNavigate } from "react-router-dom";
import { useState, useEffect } from "react";
import { ActionSync, Syncs } from "../../wailsjs/go/sync/App";

DataTable.use(DT);

function Sync(){
    const navigate = useNavigate();
    const [tableData, setTableData] = useState([]);
    
    useEffect(() => {
        Syncs().then((result:any) => {
            const jsonResult = JSON.parse(result);
            if(jsonResult.status){
                setTableData(jsonResult.data);
            }
        })
    }, []);

    const ActionUpload = () => {
        ActionSync(localStorage.getItem('user') || "").then((result:any) => {
            const jsonResult = JSON.parse(result);
            if(jsonResult.status){
                window.location.reload();
            }else{
                alert(jsonResult.message);
            }
        })
    }

    const ActionShowDetail = (event: any) => {
        const id = event.target.getAttribute('data-id');
        navigate(`/sync/upload/${id}`);
    }

    const columns = [
        { 
            data: 'start_date',
            render: (data: any) => {
                const date = new Date(data);
                return date.toLocaleString();
            }
        },
        { 
            data: 'end_date',
            render: (data: any) => {
                if(data == 0){
                    return '';
                }
                const date = new Date(data);
                return date.toLocaleString();
            }
        },
        { data: 'message' },
        { data: 'status' },
        { data: null }
    ];
    return(
        <Container fluid className="mt-3">
            <Row>
                <Col><h3>Upload Data</h3></Col>
                <Col className="d-flex justify-content-end"><Button variant="primary" onClick={ActionUpload}>Upload ke Spreadsheet</Button></Col>
            </Row>
            <DataTable data={tableData} columns={columns} 
                slots={{
                    4: (data: any, row: any) => (
                        <Button onClick={ActionShowDetail} data-id={data.id} size="sm">
                            Detail
                        </Button>
                    )
                }}
            className="table table-sm table-bordered">
                <thead>
                    <tr>
                        <th>Start Date</th>
                        <th>End Date</th>
                        <th>Message</th>
                        <th>Status</th>
                        <th>&nbsp;</th>
                    </tr>
                </thead>
            </DataTable>
        </Container>
    )
}
export default Sync