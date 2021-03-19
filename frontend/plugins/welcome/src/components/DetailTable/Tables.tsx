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
import { EntDetail } from '../../api';
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
  const [departmentid, setDepartmentid] = useState(String);
  const [details, setDetail] = useState<EntDetail[]>([])
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  //------------------------------------------------------------------------------
  const Departmentidhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setDepartmentid(event.target.value as string);
  };
  const cleardata = () => {
    setDepartmentid("");
    setStatus(false);
    setDetail([]);
  }
  //------------------------------------------------------------------------------
  const SearchDetail = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/detailtables?detail=${departmentid}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setDetail([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setDetail(data.data);
          }
        }
      });
  }
  //------------------------------------------------------------------------------
  return (
    <Page theme={pageTheme.website}>
      <Header
        title={`Department Information`}
        subtitle="ค้นหาข้อมูลแผนกแพทย์">
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
                    ค้นหาข้อมูลแผนกแพทย์
            </h1>
                </div>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>กรุณากรอกรหัสประจำแผนก</div>
                    <TextField
                      id="Departmentid"
                      value={departmentid}
                      onChange={Departmentidhandlehange || ''}
                      type="string"
                      size="medium"
                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchDetail();
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
                    แสดงข้อมูลแผนกแพทย์ {departmentid}
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่พบข้อมูลข้อมูลแผนกแพทย์
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
                  <TableCell align="left">No.</TableCell>
                  <TableCell align="left">รหัสประจำแผนก</TableCell>
                  <TableCell align="left">ชื่อแผนก</TableCell>
                  <TableCell align="left">ชื่อแพทย์หัวหน้าแผนก</TableCell>
                  <TableCell align="left">หลักสูตรอบรม</TableCell>
                  <TableCell align="left">เบอร์โทร</TableCell>
                  <TableCell align="left">อีเมล</TableCell>
                  <TableCell align="left">รายละเอียดแผนก</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {details.map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="left">{item.id}</TableCell>
                    <TableCell align="left">{item.departmentid}</TableCell>
                    <TableCell align="left">{item.edges?.Department?.Name}</TableCell>
                    <TableCell align="left">{item.edges?.Doctor?.name}</TableCell>
                    <TableCell align="left">{item.edges?.Course?.namecourse}</TableCell>
                    <TableCell align="left">{item.phone}</TableCell>
                    <TableCell align="left">{item.email}</TableCell>
                    <TableCell align="left">{item.explain}</TableCell>
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