import React, { useState } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import moment from 'moment';
import { Page, pageTheme, Header, Content } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { EntDoctor } from '../../api';
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

export default function ComponentsTable() {
  const classes = useStyles();
  const [name, setName] = useState(String);
  const [doctors, setDoctor] = useState<EntDoctor[]>([])
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  //------------------------------------------------------------------------------
  const Namehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setName(event.target.value as string);
  };
  const cleardata = () => {
    setName("");
    setStatus(false);
    setDoctor([]);
  }
  //------------------------------------------------------------------------------
  const SearchDoctor = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchpersonalinformations?doctor=${name}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setDoctor([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setDoctor(data.data);
          }
        }
      });
  }
  //------------------------------------------------------------------------------
  return (
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ค้นหาข้อมูลส่วนตัวแพทย์">
        <table>
          <Button color="secondary" variant="contained" href="/home">BACK</Button>
        </table>
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
                    ค้นหาข้อมูลส่วนตัวแพทย์
            </h1>
                </div>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>กรุณากรอกชื่อแพทย์</div>
                    <TextField
                      id="Doctoridcard"
                      value={name}
                      onChange={Namehandlehange || ''}
                      type="string"
                      size="medium"
                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchDoctor();
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
                    แสดงข้อมูลส่วนตัวแพทย์ {name}
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่พบข้อมูลข้อมูลข้อมูลส่วนตัวแพทย์
                    </Alert>
                  )}
              </div>
            ) : null}
          </Grid>
        </Grid>
        <Grid className={classes.paper}>
          <TableContainer component={Paper}>
            <Table className={classes.table} aria-label="simple table">
              <TableHead>
                <TableRow>
                <TableCell align="left">Title</TableCell>
                <TableCell align="left">Name</TableCell>
                <TableCell align="left">Position</TableCell>
                <TableCell align="left">Gender</TableCell>
                <TableCell align="left">Age</TableCell>
                <TableCell align="left">Disease</TableCell>
                <TableCell align="left">Email</TableCell>
                <TableCell align="left">Password</TableCell>
                <TableCell align="left">Phone</TableCell>
                <TableCell align="left">Address</TableCell>
                <TableCell align="left">Educational</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {doctors.map((item: any) => (
                  <TableRow key={item.id}>
                  <TableCell align="left">{item.edges?.Title?.title}</TableCell>
                  <TableCell align="left">{item.name}</TableCell>
                  <TableCell align="left">{item.edges?.Position?.position}</TableCell>
                  <TableCell align="left">{item.edges?.Gender?.gender}</TableCell>
                  <TableCell align="left">{item.age}</TableCell>
                  <TableCell align="left">{item.edges?.Disease?.disease}</TableCell>
                  <TableCell align="left">{item.email}</TableCell>
                  <TableCell align="left">{item.password}</TableCell>
                  <TableCell align="left">{item.phone}</TableCell>
                  <TableCell align="left">{item.address}</TableCell>
                  <TableCell align="left">{item.educational}</TableCell>
                    <TableCell align="left">
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </TableContainer>
        </Grid>
      </Content>
    </Page>
  );
}