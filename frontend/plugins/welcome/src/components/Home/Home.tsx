import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import { Link as RouterLink } from 'react-router-dom';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import {
  Content,
  ContentHeader,
} from '@backstage/core';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import { AppBar, Toolbar, Typography } from '@material-ui/core';

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    justifyContent: 'center'
  },
  paper: {
    height: 450,
    width: 800,
    backgroundImage: 'url(https://www.okusanpix.com/wp-content/uploads/2019/09/3548811765a393c9a3cd79211e8bf6d8-800x450.png)',
  },
  control: {
    padding: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },

})
);

const Homepage: FC<{}> = () => {
  const classes = useStyles();


  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบข้อมูลแพทย์
          </Typography>
          <Button color="inherit" component={RouterLink} to="/" startIcon={<ExitToAppRoundedIcon />}> Logout </Button>
        </Toolbar>
      </AppBar>

      <Content>
        <ContentHeader title="ระบบข้อมูลแพทย์"> </ContentHeader>
        <Breadcrumbs aria-label="breadcrumb" >
          <Link
            color="textPrimary"
            href="/history" >
            บันทึกประวัติการทำงานของแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/training" >
            บันทึกข้อมูลการฝึกอบรม
            </Link>

          <Link
            color="textPrimary"
            href="/detail" >
            บันทึกข้อมูลแผนกแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/schedule" >
            บันทึกเวลาของแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/specialdoctor" >
            บันทึกข้อมูลแพทย์เฉพาะทาง
            </Link>
        </Breadcrumbs>
        <br></br>
        <h1>ค้นหา</h1>
        <Breadcrumbs>
          <Link
            color="textPrimary"
            href="/workhistorytable" >
            ค้นหาประวัติการทำงานของแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/trainingTable" >
            ค้นหาข้อมูลการฝึกอบรม
            </Link>

          <Link
            color="textPrimary"
            href="/detailtable" >
            ค้นหาข้อมูลแผนกแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/personalinformationtable" >
            ค้นหาข้อมูลส่วนตัวแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/scheduletable" >
            ค้นหาเวลาของแพทย์
            </Link>

          <Link
            color="textPrimary"
            href="/specialdoctortable" >
            ค้นหาข้อมูลแพทย์เฉพาะทาง
          </Link>

        </Breadcrumbs>
        <br></br>
        <Grid container className={classes.root} spacing={2}>
          <Grid item xs={12}>
            <Grid container justify="center" >
              {[0].map((value) => (
                <Grid key={value} item>
                  <Paper className={classes.paper} />
                </Grid>
              ))}
            </Grid>
          </Grid>
          <Grid item xs={12}>
          </Grid>
        </Grid>
      </Content>


    </div>
  );
};
export default Homepage;