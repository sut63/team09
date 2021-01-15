import React, { useEffect, FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Content } from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Grid from '@material-ui/core/Grid';
import { DefaultApi } from '../../api/apis';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Select from '@material-ui/core/Select';
import { InputLabel, MenuItem } from '@material-ui/core';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntDepartment } from '../../api/models/EntDepartment';
import { EntExtradoctor } from '../../api/models/EntExtradoctor';
import Swal from 'sweetalert2';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1),
      minWidth: 400,
      maxWidth: 400,
      marginTop: '1.5%',
    },
    withoutLabel: {
      marginTop: theme.spacing(4),
    },
    title: {
      flexGrow: 1,
    },

  }),
);
interface specialdoctor {
  extradoctor: number;
  doctor: number;
  other: string;
  department: number;
}

const Specialdoctor: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [doctors, setDoctor] = React.useState<EntDoctor[]>([]);
  const [departments, setDepartment] = React.useState<EntDepartment[]>([]);
  const [extradoctors, setExtradoctor] = React.useState<EntExtradoctor[]>([]);

  const [specialdoctor, setSpecialdoctor] = React.useState<
    Partial<specialdoctor>>({});

  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 1500
  });
  //set data object path
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Specialdoctor;
    const { value } = event.target;
    setSpecialdoctor({ ...specialdoctor, [name]: value });
  };
  const handleChange2 = (
    event: React.ChangeEvent<{ name: string; value: number }>,
  ) => {
    const name = event.target.name as keyof typeof Specialdoctor;
    const { value } = event.target;
    setSpecialdoctor({ ...specialdoctor, [name]: +value });
  };
  console.log(specialdoctor);
  const getDoctors = async () => {
    const res = await http.listDoctor({ limit: 10, offset: 0 });
    setDoctor(res);
    console.log(res)
  };

  const getDepartments = async () => {
    const res = await http.listDepartment({ limit: 10, offset: 0 });
    setDepartment(res);
    console.log(res)
  };

  const getExtradoctors = async () => {
    const res = await http.listExtradoctor({ limit: 11, offset: 0 });
    setExtradoctor(res);
    console.log(res)
  };

  // Lifecycle Hooks
  useEffect(() => {
    getDoctors();
    getDepartments();
    getExtradoctors();
    getDoctors();
    getDepartments();

  }, []);

  function clear() {
    setSpecialdoctor({});
  }

  function save() {
    if (specialdoctor.department == null || specialdoctor.doctor == null || specialdoctor.other == null || specialdoctor.extradoctor == null) {
      Toast.fire({
        title: 'บันทึกข้อมูลไม่สำเร็จ',
        icon: 'error',
      });
    } else {
      Toast.fire({
        title: 'บันทึกข้อมูลสำเร็จ',
        icon: 'success',
      });
    }
    const apiUrl = 'http://localhost:8080/api/v1/specialdoctors';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(specialdoctor),
    };
    console.log(specialdoctor); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            title: 'บันทึกข้อมูลสำเร็จ',
            icon: 'success',
          });
        } else {
          Toast.fire({
            title: 'บันทึกข้อมูลไม่สำเร็จ',
            icon: 'error',
          });
        }
      });
  }
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
      <Content className={classes.withoutLabel}>
        <Typography variant="h5"
          className={classes.formControl} style={{ marginLeft: 100 }}>
          บันทึกข้อมูลแพทย์เฉพาะทาง
        </Typography>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>ชื่อ-นามสกุล</InputLabel>
                <Select
                  name="doctor"
                  label="ชื่อ-นามสกุล"
                  type="string"

                  value={specialdoctor.doctor || ''}
                  onChange={handleChange}
                >
                  {doctors.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>แผนก</InputLabel>
                <Select
                  name="department"
                  type="string"
                  label="แผนก"
                  value={specialdoctor.department || ''}
                  onChange={handleChange}
                >
                  {departments.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>สาขาเฉพาะทาง</InputLabel>
                <Select
                  name="extradoctor"
                  label="สาขาเฉพาะทาง"
                  type="string"
                  value={specialdoctor.extradoctor || ''}
                  onChange={handleChange}
                >
                  {extradoctors.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.specialname}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
              >
                <TextField
                  name="other"
                  label="หมายเหตุ"
                  type="string"
                  variant="outlined"
                  value={specialdoctor.other || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <div className={classes.formControl} style={{ marginLeft: 180 }}>
              <Button
                onClick={save}
                variant="contained"
                color="primary">
                SAVE
             </Button>
              <Button style={{ marginLeft: 10 }}
                onClick={clear}
                variant="contained"
                color="secondary">
                CLEAR
             </Button>
              <Button style={{ marginLeft: 10 }}
                component={RouterLink} to="/home"
                variant="contained"
                color="secondary">
                BACK
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </div>
  );
};
export default Specialdoctor;