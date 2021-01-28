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
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntDoctor } from '../../api/models/EntDoctor';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import { InputBase, Link } from '@material-ui/core';
import moment from 'moment';
import SearchIcon from '@material-ui/icons/Search';
import TextField from '@material-ui/core/TextField';


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
    search: {
      position: 'relative',
      borderRadius: theme.shape.borderRadius,
      backgroundColor: fade(theme.palette.common.black, 0.15),
      '&:hover': {
        backgroundColor: fade(theme.palette.common.black, 0.25),
      },
      marginLeft: 0,
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        marginLeft: theme.spacing(1),
        width: 'auto',
      },
    },
    searchIcon: {
      padding: theme.spacing(0, 2),
      height: '100%',
      position: 'absolute',
      pointerEvents: 'none',
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'center',
    },
    inputRoot: {
      color: 'inherit',
    },
    inputInput: {
      padding: theme.spacing(1, 1, 1, 0),
      // vertical padding + font size from searchIcon
      paddingLeft: `calc(1em + ${theme.spacing(4)}px)`,
      transition: theme.transitions.create('width'),
      width: '100%',
      [theme.breakpoints.up('sm')]: {
        width: '12ch',
        '&:focus': {
          width: '20ch',
        },
      },
    },
  }),
);

export default function ComponentsTableUser() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [doctors, setDoctor] = useState<EntDoctor[]>([]);
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = useState('');

  useEffect(() => {
    const getDoctor = async () => {
      const res = await http.listDoctor({ limit: 10, offset: 0 });
      setLoading(false);
      setDoctor(res);
    };
    getDoctor();
  }, [loading]);

  const deleteDoctors = async (id: number) => {
    const res = await http.deleteDoctor({ id: id });
    setLoading(true);
  };

 
  const filterDoctor = doctors.filter( doctors => {
    return doctors.name?.includes(search)
  })


  function emptyDoctor () : any { 
    if(filterDoctor.length == 0){
      const doctord = <TableRow> <TableCell align="center" colSpan={9}><p>ไม่มีข้อมูลในระบบ</p></TableCell></TableRow>;
      return doctord;
    }
  }
  
  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ข้อมูลส่วนตัวแพทย์"
      ></Header>
      <Content>
      <ContentHeader title="ข้อมูลส่วนตัวแพทย์">
      <div className={classes.search} style={{ marginRight: 10 }}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <InputBase
              style={{ marginRight: 20 }}
              placeholder="Search…"
              value = {search} 
              onChange = {(event) => {setSearch(event.target.value);}}
              classes={{
                root: classes.inputRoot,
                input: classes.inputInput,
              }}
              inputProps={{ 'aria-label': 'search' }}
            />
          </div>
        <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
           บันทึกข้อมูลส่วนตัวแพทย์
           </Button>
         </Link>
       </ContentHeader>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">Title</TableCell>
                <TableCell align="center">Name</TableCell>
                <TableCell align="center">Position</TableCell>
                <TableCell align="center">Gender</TableCell>
                <TableCell align="center">Age</TableCell>
                <TableCell align="center">Disease</TableCell>
                <TableCell align="center">Email</TableCell>
                <TableCell align="center">Password</TableCell>
                <TableCell align="center">Phone</TableCell>
                <TableCell align="center">Address</TableCell>
                <TableCell align="center">Educational</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>

              {emptyDoctor()}
              {filterDoctor.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.title?.title}</TableCell>
                  <TableCell align="center">{item.name}</TableCell>
                  <TableCell align="center">{item.edges?.position?.position}</TableCell>
                  <TableCell align="center">{item.edges?.gender?.gender}</TableCell>
                  <TableCell align="center">{item.age}</TableCell>
                  <TableCell align="center">{item.edges?.disease?.disease}</TableCell>
                  <TableCell align="center">{item.email}</TableCell>
                  <TableCell align="center">{item.password}</TableCell>
                  <TableCell align="center">{item.phone}</TableCell>
                  <TableCell align="center">{item.address}</TableCell>
                  <TableCell align="center">{item.educational}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                       if (item.id === undefined) {
                          return;
                        }
                        deleteDoctors(item.id);
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