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
import { EntUser } from '../../api/models/EntUser';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import { Link } from '@material-ui/core';

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
  const [users, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getUsers = async () => {
      const res = await http.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUsers();
  }, [loading]);

  const deleteUsers = async (id: number) => {
    const res = await http.deleteUser({ id: id });
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
         <Link component={RouterLink} to="/user">
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
                <TableCell align="center">Specialist Doctor</TableCell>
                <TableCell align="center">Office</TableCell>
                <TableCell align="center">Time Period</TableCell>

              </TableRow>
            </TableHead>
            <TableBody>
              {users.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.title?.title}</TableCell>
                  <TableCell align="center">{item.name}</TableCell>
                  <TableCell align="center">{item.edges?.position?.position}</TableCell>
                  <TableCell align="center">{item.edges?.gender?.gender}</TableCell>
                  <TableCell align="center">{item.email}</TableCell>
                  <TableCell align="center">{item.password}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                       if (item.id === undefined) {
                          return;
                        }
                        deleteUsers(item.id);
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
