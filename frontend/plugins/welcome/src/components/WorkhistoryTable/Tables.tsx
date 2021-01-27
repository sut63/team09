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
import { EntOffice } from '../../api/models/EntOffice';
import { Content, ContentHeader, Header, Page } from '@backstage/core';
import { pageTheme } from '@backstage/core';
import { Link } from '@material-ui/core';
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
        width: '30ch',
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
  const [offices, setOffice] = React.useState<EntOffice[]>([]);
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = React.useState('');

  useEffect(() => {
    const getOffice = async () => {
      const res = await http.listOffice({ limit: 10, offset: 0 });
      setLoading(false);
      setOffice(res);
    };
    getOffice();
  }, [loading]);

  const deleteOffices = async (id: number) => {
    const res = await http.deleteOffice({ id: id });
    setLoading(true);
  };

  const filterOffice = offices.filter(offices => {
    return offices.edges?.doctor?.name?.includes(search)
  })


  function emptyOffice () : any { 
    if(filterOffice.length == 0){
      const officed = <TableRow> <TableCell align="center" colSpan={9}><p>ไม่มีข้อมูลในระบบ</p></TableCell></TableRow>;
      return officed;
    }
  }

  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ประวัติการทำงานของแพทย์">
      </Header>
      <Content>
        <ContentHeader title="ประวัติการทำงานของแพทย์">
          <div className={classes.search} style={{ marginRight: 10 }}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <TextField
              name="search"
              className={classes.inputInput}
              style={{ marginRight: 100 }}
              placeholder="Search"
              value={search}
              onChange={(event) => { setSearch(event.target.value); }}
              inputProps={{ 'aria-label': 'search' }}
            />
          </div>
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
                <TableCell align="center">Firsttime</TableCell>
                <TableCell align="center">Finallytime</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>

              {/* ต้องใช้เงื่อนไข check ว่า filterOffice.length == 0 */}
              {emptyOffice()}
              {filterOffice.map(item => (
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
