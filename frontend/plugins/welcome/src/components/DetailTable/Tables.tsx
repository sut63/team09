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
import { EntDetail } from '../../api/models/EntDetail';
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
  const [details, setDetail] = React.useState<EntDetail[]>([]);
  const [loading, setLoading] = useState(false);
  const [search, setSearch] = React.useState('');

  useEffect(() => {
    const getDetail = async () => {
      const res = await http.listDetail({ limit: 10, offset: 0 });
      setLoading(false);
      setDetail(res);
    };
    getDetail();
  }, [loading]);

  const deleteDetails = async (id: number) => {
    const res = await http.deleteDetail({ id: id });
    setLoading(true);
  };

  const filterDetail = details.filter(details => {
    return details.edges?.department?.name?.includes(search)
  })


  function emptyDetail () : any { 
    if(filterDetail.length == 0){
      const detailed = <TableRow> <TableCell align="center" colSpan={9}><p>ไม่มีข้อมูลในระบบ</p></TableCell></TableRow>;
      return detailed;
    }
  }

  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ข้อมูลแผนกแพทย์">
        <Button color="secondary" variant="contained" href="/home">BACK</Button>
      </Header>
      <Content>
        <ContentHeader title="ข้อมูลแผนกแพทย์">
          <div className={classes.search} style={{ marginRight: 10 }}>
            <div className={classes.searchIcon}>
              <SearchIcon />
            </div>
            <TextField
              name="search"
              className={classes.inputInput}
              style={{ marginRight: 100 }}
              placeholder="กรุณากรอกชื่อแผนก"
              value={search}
              onChange={(event) => { setSearch(event.target.value); }}
              inputProps={{ 'aria-label': 'search' }}
            />
          </div>
          <Link component={RouterLink} to="/detail">
            <Button variant="contained" color="primary">
              บันทึกข้อมูลแผนกแพทย์
           </Button>
          </Link>
        </ContentHeader>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
              <TableCell align="center">Department</TableCell>
                <TableCell align="center">Mission</TableCell>
                <TableCell align="center">Course</TableCell>
                <TableCell align="center">Phone</TableCell>
                <TableCell align="center">Email</TableCell>
                <TableCell align="center">Detail</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>

              {/* ต้องใช้เงื่อนไข check ว่า filterOffice.length == 0 */}
              {emptyDetail()}
              {filterDetail.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.edges?.department?.name}</TableCell>
                  <TableCell align="center">{item.edges?.mission?.mission}</TableCell>
                  <TableCell align="center">{item.edges?.course?.namecourse}</TableCell>
                  <TableCell align="center">{item.phone}</TableCell>
                  <TableCell align="center">{item.email}</TableCell>
                  <TableCell align="center">{item.explain}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                        if (item.id === undefined) {
                          return;
                        }
                        deleteDetails(item.id);
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
