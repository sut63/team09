import React, { useState } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Page, pageTheme, Header, Content } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntSpecialdoctor } from '../../api';
import { Alert } from '@material-ui/lab';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch: {
      width: 'auto',
      margin: '10px',
      color: '#FFFFFF',
      background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },
  }),
);
 //------------------------------------------------------------------------------
export default function ComponentsTableUser() {
  const classes = useStyles();
  const [doctorid, setDoctorid] = useState(String);
  const [specialdoctors, setSpecialdoctor] = useState<EntSpecialdoctor[]>([]);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);
  
  //------------------------------------------------------------------------------

  const Doctoridhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setDoctorid(event.target.value as string);
  };
  const cleardata = () => {
    setDoctorid("");
    setStatus(false);
    setSpecialdoctor([]);
  }
  //------------------------------------------------------------------------------
  const SearchSpecialdoctor = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchspecialdoctors?specialdoctor=${doctorid}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setSpecialdoctor([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setSpecialdoctor(data.data);
          }
        }
      });
  }
  //------------------------------------------------------------------------------
 
  return (
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ค้นหาข้อมูลแพทย์เฉพาะทาง">
        <Button color="secondary" variant="contained" href="/home">BACK</Button>
      </Header>
      <Content>
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>
              <Typography align="center" >
                <div style={{ background: "#3792cb", height: 50 }}>
                  <h1 style={
                    {
                      color: "#FFFFFF",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '30px',
                    }}>
                    ค้นหาข้อมูลแพทย์เฉพาะทาง
            </h1>
                </div>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>กรุณากรอกรหัสประจำตัวแพทย์</div>
                    <TextField
                      id="Doctorid"
                      value={doctorid}
                      onChange={Doctoridhandlehange || ''}
                      type="string"
                      size="medium"
                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchSpecialdoctor();
                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#3792cb", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',
                      }
                    }>
                    ค้นหาข้อมูล
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();
                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#3792cb", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 25px',
                      }
                    }>
                    เคลียร์ข้อมูล
            </h3>
                </Button>
              </Typography>
            </Paper>
            {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    แสดงข้อมูลแพทย์เฉพาะทาง {doctorid}
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่พบข้อมูลข้อมูลประวัติการทำงานของแพทย์
                    </Alert>
                  )}
              </div>
            ) : null}
          </Grid>
        </Grid>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">No.</TableCell>
                  <TableCell align="center">Name</TableCell>
                  <TableCell align="center">Department</TableCell>
                  <TableCell align="center">Specialist</TableCell>
                  <TableCell align="center">DoctorIDCard</TableCell>
                  <TableCell align="center">Roomname</TableCell>
                <TableCell align="center">Other</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
            {specialdoctors.map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                    <TableCell align="center">{item.edges?.Doctor?.name}</TableCell>
                    <TableCell align="center">{item.edges?.Department?.Name}</TableCell>
                    <TableCell align="center">{item.edges?.Extradoctor?.specialname}</TableCell>
                    <TableCell align="center">{item.Doctorid}</TableCell>
                    <TableCell align="center">{item.Roomnumber}</TableCell>
                    <TableCell align="center">{item.Other}</TableCell>
                    <TableCell align="center">
                    </TableCell>
                  </TableRow>
                ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </Page>

  );
}
