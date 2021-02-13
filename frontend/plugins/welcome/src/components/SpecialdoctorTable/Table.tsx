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
import { EntSpecialdoctor } from '../../api/models/EntSpecialdoctor';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import { InputBase, Link, TextField } from '@material-ui/core';
import moment from 'moment';
import SearchIcon from '@material-ui/icons/Search';


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
        width: '30ch',
        '&:focus': {
          width: '40ch',
        },
      },
    },
  }),
);



export default function ComponentsTableUser() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [specialdoctors, setSpecialdoctor] = React.useState<EntSpecialdoctor[]>([]);
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = useState('');
  
  const getSpecialdoctor = async () => {
    const res = await http.listSpecialdoctor({ limit: 10, offset: 0 });
    setLoading(false);
    setSpecialdoctor(res);
    console.log(res);
  };
  useEffect(() => {
    getSpecialdoctor();
    
  }, [loading]);

  const deleteSpecialdoctors = async (id: number) => {
    const res = await http.deleteSpecialdoctor({ id: id });
    
    setLoading(true);
  };
  // useEffect(() => {
  //   setfilterSpecialdoctor (specialdoctors.filter( specialdoctors => {
  //     return specialdoctors.edges?.doctor?.name?.includes(search)
  //   })
  //   )
  // },[se])

  const filterspecialdoctor = specialdoctors.filter( specialdoctors => {
    return specialdoctors.edges?.doctor?.name?.includes(search)
  })
  function emptySpecialdoctor () : any { 
    if(filterspecialdoctor.length == 0){
      const special = <TableRow> <TableCell align="center" colSpan={9}><p>ไม่มีข้อมูลในระบบ</p></TableCell></TableRow>;
      return special;
    }
  }

  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ข้อมูลแพทย์เฉพาะทาง">
        <Button color="secondary" variant="contained" href="/home">BACK</Button>
      </Header>
      <Content>
        <ContentHeader title="ข้อมูลแพทย์เฉพาะทาง">
          <div className={classes.search} style={{ marginRight: 10 }}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <TextField
              name="search"
              className={classes.inputInput}
              style={{ marginRight: 100 }}
              placeholder="กรุณากรอกชื่อแพทย์"
              value={search}
              onChange={(event) => { setSearch(event.target.value); }}
              inputProps={{ 'aria-label': 'search' }}
            />
          </div>
          <Link component={RouterLink} to="/specialdoctor">
            <Button variant="contained" color="primary">
              บันทึกข้อมูลแพทย์เฉพาะทาง
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
                <TableCell align="center">DoctorIDCard</TableCell>
                <TableCell align="center">Roomnumber</TableCell>
                <TableCell align="center">Other</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
            {emptySpecialdoctor()}
              {filterspecialdoctor.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.doctor?.name}</TableCell>
                  <TableCell align="center">{item.edges?.department?.name}</TableCell>
                  <TableCell align="center">{item.edges?.extradoctor?.specialname}</TableCell>
                  <TableCell align="center">{item.doctorid}</TableCell>
                  <TableCell align="center">{item.roomnumber}</TableCell>
                  <TableCell align="center">{item.other}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        if (item.id === undefined) {
                          return;
                        }
                        deleteSpecialdoctors(item.id);
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
