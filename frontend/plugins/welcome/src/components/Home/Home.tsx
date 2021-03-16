import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import clsx from 'clsx';
import { makeStyles, useTheme, Theme, createStyles } from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import CssBaseline from '@material-ui/core/CssBaseline';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';
import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import Button from '@material-ui/core/Button';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import { Grid, Paper } from '@material-ui/core';


const drawerWidth = 240;

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
    },
    appBar: {
      transition: theme.transitions.create(['margin', 'width'], {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
    },
    appBarShift: {
      width: `calc(100% - ${drawerWidth}px)`,
      marginLeft: drawerWidth,
      transition: theme.transitions.create(['margin', 'width'], {
        easing: theme.transitions.easing.easeOut,
        duration: theme.transitions.duration.enteringScreen,
      }),
    },
    menuButton: {
      marginRight: theme.spacing(2),
    },
    hide: {
      display: 'none',
    },
    drawer: {
      width: drawerWidth,
      flexShrink: 0,
    },
    drawerPaper: {
      width: drawerWidth,
    },
    drawerHeader: {
      display: 'flex',
      alignItems: 'center',
      padding: theme.spacing(0, 1),
      // necessary for content to be below app bar
      ...theme.mixins.toolbar,
      justifyContent: 'flex-end',
    },
    content: {
      flexGrow: 1,
      padding: theme.spacing(3),
      transition: theme.transitions.create('margin', {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      marginLeft: -drawerWidth,
    },
    contentShift: {
      transition: theme.transitions.create('margin', {
        easing: theme.transitions.easing.easeOut,
        duration: theme.transitions.duration.enteringScreen,
      }),
      marginLeft: 0,
    },
    title: {
      flexGrow: 1,
    },
    paper: {
      height: 500,
      width: 1200,
      marginTop: 90,
      backgroundImage: 'url(https://www.img.in.th/images/fb357cfcfd8ebe7bcafddec82ed50856.jpg)',
    },
  }),
);

export default function PersistentDrawerLeft() {
  const classes = useStyles();
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  return (
    <div className={classes.root}>
      <CssBaseline />
      <AppBar
        position="fixed"
        className={clsx(classes.appBar, {
          [classes.appBarShift]: open,
        })}
      >
        <Toolbar>
          <IconButton
            color="inherit"
            aria-label="open drawer"
            onClick={handleDrawerOpen}
            className={clsx(classes.menuButton, open && classes.hide)}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            ระบบข้อมูลแพทย์
          </Typography>
          <Button color="inherit" component={RouterLink} to="/" startIcon={<ExitToAppRoundedIcon />}> Logout </Button>
        </Toolbar>
      </AppBar>
      <Drawer
        className={classes.drawer}
        variant="persistent"
        anchor="left"
        open={open}
        classes={{
          paper: classes.drawerPaper,
        }}
      >
        <div className={classes.drawerHeader}>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === 'ltr' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
          </IconButton>
        </div>
        <Divider />
        <List>
          
          <Button color="inherit" component={RouterLink} to="/training">บันทึกข้อมูลการฝึกอบรม</Button>
          <Button color="inherit" component={RouterLink} to="/detail">บันทึกข้อมูลแผนกแพทย์</Button>
          <Button color="inherit" component={RouterLink} to="/schedule">บันทึกเวลาของแพทย์</Button>
          <Button color="inherit" component={RouterLink} to="/specialdoctor">บันทึกข้อมูลแพทย์เฉพาะทาง</Button>
          <Button color="inherit" component={RouterLink} to="/history">บันทึกประวัติการทำงานของแพทย์</Button>
        </List>
        <Divider />
        <List>
          <Button color="inherit" component={RouterLink} to="/trainingtable">ค้นหาการฝึกอบรม</Button>
          <Button color="inherit" component={RouterLink} to="/detailtable">ค้นหาข้อมูลแผนกแพทย์</Button>
          <Button color="inherit" component={RouterLink} to="/scheduletable">ค้นหาตารางเวลาของแพทย์</Button>
          <Button color="inherit" component={RouterLink} to="/personalinformationtable">ค้นหาข้อมูลส่วนตัวแพทย์</Button>
          <Button color="inherit" component={RouterLink} to="/specialdoctortable">ค้นหาข้อมูลแพทย์เฉพาะทาง</Button>
          <Button color="inherit" component={RouterLink} to="/searchworkhistory">ค้นหาประวัติการทำงานของแพทย์</Button>
        </List>
        <Divider />
      </Drawer>
      <main
        className={clsx(classes.content, {
          [classes.contentShift]: open,
        })}
      >
      <div className={classes.drawerHeader}/>
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
        </main>
    </div>
  );
}