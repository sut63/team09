import React, { useState, useEffect } from 'react';
import { createStyles, fade, makeStyles, Theme } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntSchedule } from '../../api/models/EntSchedule';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import moment from 'moment';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';






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



export default function ComponentsTableUser() {
  const classes = useStyles();
  const [docterid, setDocterid] = useState(String);
  const [schedules, setSchedule] = useState<EntSchedule[]>([]);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  ///////////////////////////////////////////////////////////////////////////////

  const Docteridhandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setDocterid(event.target.value as string);
  };
  const cleardata = () => {
    setDocterid("");
    setStatus(false);
    setSchedule([]);
  }
  ///////////////////////////////////////////////////////////////////////////////

  const SearchSchedule = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchschedules?schedule=${docterid}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setSchedule([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setSchedule(data.data);
          }
        }
      });
  }
  ///////////////////////////////////////////////////////////////////////////////


  return (
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ค้นหาตารางเวลาของแพทย์">
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
                    ค้นหาตารางเวลาของแพทย์
            </h1>
                </div>
                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>กรุณากรอกรหัสประจำตัวแพทย์</div>
                    <TextField
                      id="Docterid"
                      value={docterid}
                      onChange={Docteridhandlehange || ''}
                      type="string"
                      size="medium"
                      style={{ width: 300 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchSchedule();
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
                    แสดงข้อมูลตารางเวลาของแพทย์ {docterid}
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่พบข้อมูลข้อมูลตารางเวลาของแพทย์
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
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">ชื่อแพทย์</TableCell>
                <TableCell align="center">แผนกของแพทย์</TableCell>
                <TableCell align="center">กิจกรรมของแพทย์</TableCell>
                <TableCell align="center">หมายเลขห้องทำงานของแพทย์</TableCell>
                <TableCell align="center">เลขประจำตัวแพทย์</TableCell>
                <TableCell align="center">สถานที่ทำงาน</TableCell>
                <TableCell align="center">วันที่และเวลา</TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {schedules.map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.Doctor?.name}</TableCell>
                  <TableCell align="center">{item.edges?.Department?.Name}</TableCell>
                  <TableCell align="center">{item.Activity}</TableCell>
                  <TableCell align="center">{item.Roomnumber}</TableCell>
                  <TableCell align="center">{item.Docterid}</TableCell>
                  <TableCell align="center">{item.edges?.Office?.officename}</TableCell>
                  <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
                    <TableCell align="center">
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