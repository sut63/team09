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
import { EntTraining } from '../../api/models/EntTraining';
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
          width: '20ch',
        },
      },
    },
  }),
);

export default function ComponentsTableUser() {
  const classes = useStyles();
  const http = new DefaultApi();
  const [trainings, setTraining] = React.useState<EntTraining[]>([]);
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState('');

  useEffect(() => {
    const getTraining = async () => {
      const res = await http.listTraining({ limit: 20, offset: 0 });
      setLoading(false);
      setTraining(res);
    };
    getTraining();
  }, [loading]);

  const deleteTrainings = async (id: number) => {
    const res = await http.deleteTraining({ id: id });
    setLoading(true);
  };

  const filterTraining = trainings.filter( trainings => {
    return trainings.edges?.doctor?.name?.includes(search)
  })

  function emptyTraining () : any { 
    if(filterTraining.length == 0){
      const training = <TableRow> <TableCell align="center" colSpan={9}><p>ไม่มีข้อมูลในระบบ</p></TableCell></TableRow>;
      return training;
    }
  }


  return (
    // <div className={classes.root}>
    <Page theme={pageTheme.website}>
      <Header
        title={`Doctor Information`}
        subtitle="ข้อมูลการฝึกอบรม">
        <Button color="secondary" variant="contained" href="/home">BACK</Button>
      </Header>

      <Content>
      <ContentHeader title="ข้อมูลการฝึกอบรม">


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


        <Link component={RouterLink} to="/training">
           <Button variant="contained" color="primary">
           บันทึกข้อมูลการฝึกอบรม
           </Button>
         </Link>
       </ContentHeader>

        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">No.</TableCell>
                <TableCell align="center">หลักสูตร</TableCell>
                <TableCell align="center">สาขา</TableCell>
                <TableCell align="center">วันแรกที่เข้าร่วมอบรม</TableCell>
                <TableCell align="center">วันสุดท้ายที่เข้าร่วมอบม</TableCell>
                <TableCell align="center">ชั่วโมงการอบรม</TableCell>
                <TableCell align="center">เลขบัตรประจำตัวแพทย์</TableCell>
                <TableCell align="center">ชื่อแพทย์</TableCell>
                <TableCell align="center">แผนก</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>

              {emptyTraining()}
              {filterTraining.map(item => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.course?.namecourse}</TableCell>
                  <TableCell align="center">{item.branch}</TableCell>
                  <TableCell align="center">{moment(item.dateone).format('DD/MM/YYYY')}</TableCell>
                  <TableCell align="center">{moment(item.datetwo).format('DD/MM/YYYY')}</TableCell>
                  <TableCell align="center">{item.hour}</TableCell>
                  <TableCell align="center">{item.doctoridcard}</TableCell>
                  <TableCell align="center">{item.edges?.doctor?.name}</TableCell>
                  <TableCell align="center">{item.edges?.department?.name}</TableCell>
                  <TableCell align="center">
                    <Button
                      onClick={() => {
                       if (item.id === undefined) {
                          return;
                        }
                        deleteTrainings(item.id);
                      }}
                      style={{ marginLeft: 5 }}
                      variant="contained"
                      color="secondary"
                    >
                      ลบ
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
