import React, { useState, useEffect } from 'react';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import { Link as RouterLink } from 'react-router-dom';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntOffice } from '../../api/models/EntOffice';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import { Link } from '@material-ui/core';
import moment from 'moment';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    table: {
      minWidth: 650,
    },
  }),
);

export default function ComponentsTableUser() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [offices, setOffice] = React.useState<EntOffice[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getOffice = async () => {
      const res = await http.listOffice({ limit: 20, offset: 0 });
      setLoading(false);
      setOffice(res);
    };
    getOffice();
  }, [loading]);

  const deleteOffices = async (id: number) => {
    const res = await http.deleteOffice({ id: id });
    setLoading(true);
  };

  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ประวัติการทำงานของแพทย์"
      ></Header>
      <Content>
      <ContentHeader title="ประวัติการทำงานของแพทย์">
        <Link component={RouterLink} to="/history">
           <Button variant="contained" color="primary">
           บันทึกประวัติการทำงานของแพทย์
           </Button>
         </Link>
       </ContentHeader>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">Name</TableCell>
                <TableCell align="center">Department</TableCell>
                <TableCell align="center">Specialist</TableCell>
                <TableCell align="center">Officename</TableCell>
                <TableCell align="center">DoctorIDCard</TableCell>
                <TableCell align="center">Roomname</TableCell>
                <TableCell align="center">First</TableCell>
                <TableCell align="center">Finally</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {offices.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.doctor?.name}</TableCell>
                  <TableCell align="center">{item.edges?.department?.name}</TableCell>
                  <TableCell align="center">{item.edges?.extradoctor?.specialname}</TableCell>
                  <TableCell align="center">{item.officename}</TableCell>
                  <TableCell align="center">{item.doctoridcard}</TableCell>
                  <TableCell align="center">{item.roomnumber}</TableCell>
                  <TableCell align="center">{moment(item.addedTime1).format('DD/MM/YYYY')}</TableCell>
                  <TableCell align="center">{moment(item.addedTime2).format('DD/MM/YYYY')}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                       if (item.id === undefined) {
                          return;
                        }
                        deleteOffices(item.id);
                      }}
                      style={{ marginLeft: 10 }}
                      variant="contained"
                      color="secondary"
                    >
                      Delete
               </Button>
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
