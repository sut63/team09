import React, { useEffect, FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
} from '@backstage/core';
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
import { IconButton, InputAdornment, InputLabel, MenuItem, OutlinedInput } from '@material-ui/core';
import { EntTitle } from '../../api/models/EntTitle';
import { EntPosition } from '../../api/models/EntPosition';
import { EntGender } from '../../api/models/EntGender';
import Swal from 'sweetalert2';
import { EntDisease } from '../../api/models/EntDisease';

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

interface doctor {
  title: number;
  gender: number;
  position: number;
  disease: number;
  name: string;
  email: string;
  educational: string;
  address: string;
  phone: string;
  age: number;
  password: string;
  showPassword: boolean;
}

const Doctor: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [doctor, setDoctor] = React.useState<Partial<doctor>>({});
  const [titles, setTitles] = React.useState<EntTitle[]>([]);
  const [positions, setPositions] = React.useState<EntPosition[]>([]);
  const [genders, setGenders] = React.useState<EntGender[]>([]);
  const [diseases, setDiseases] = React.useState<EntDisease[]>([]);


  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 1500
  });

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
    const name = event.target.name as keyof typeof Doctor;
    const { value } = event.target;
    setDoctor({ ...doctor, [name]: value });
    console.log(doctor);
  };

  const handleChangeNum = (
    event: React.ChangeEvent<{ name: string; value: number }>,) => {
    const name = event.target.name as keyof typeof Doctor;
    const { value } = event.target;
    setDoctor({ ...doctor, [name]: +value });
    // console.log(Doctor);
  };

  const getTitles = async () => {
    const res = await http.listTitle({ limit: 10, offset: 0 });
    setTitles(res);
  };
  const getGenders = async () => {
    const res = await http.listGender({ limit: 10, offset: 0 });
    setGenders(res);
  };
  const getPositions = async () => {
    const res = await http.listPosition({ limit: 10, offset: 0 });
    setPositions(res);
  };
  const getDiseases = async () => {
    const res = await http.listDisease({ limit: 10, offset: 0 });
    setDiseases(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getGenders();
    getPositions();
    getTitles();
    getDiseases();
  }, []);

  function clear() {
    setDoctor({});
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/doctors';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(doctor),
    };
    console.log(doctor); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data); 
        if (data.status == true) {
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
          <Typography variant="h4" className={classes.title}>
            ระบบข้อมูลแพทย์
            </Typography>
          <Button color="inherit" component={RouterLink} to="/"> Login </Button>
        </Toolbar>
      </AppBar>
      <Content className={classes.withoutLabel}>
        <Typography variant="h3"
          className={classes.formControl} style={{ marginLeft: 120 }}>
          บันทึกข้อมูลส่วนตัวแพทย์
        </Typography>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={3}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>คำนำหน้าชื่อ</InputLabel>
                <Select
                  name="title"
                  label="คำนำหน้าชื่อ"
                  value={doctor.title || ''}
                  onChange={handleChange} >
                  {titles.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.title}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={3}>

              <Grid item xs={3}>
                <FormControl
                  fullWidth
                  className={classes.formControl}
                  style={{ marginLeft: 100 }}
                  variant="outlined">
                  <TextField
                    name="name"
                    label="ชื่อ-สกุล"
                    variant="outlined"
                    size="medium"
                    value={doctor.name || ''}
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>           
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>ตำแหน่ง</InputLabel>
                <Select
                  name="position"
                  type="string"
                  label="ตำแหน่ง"
                  value={doctor.position || ''}
                  onChange={handleChange}
                >
                  {positions.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.position}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>เพศ</InputLabel>
                <Select
                  name="gender"
                  label="เพศ"
                  type="string"
                  value={doctor.gender || ''}
                  onChange={handleChange}
                >
                  {genders.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.gender}</MenuItem>
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
                variant="outlined">
                <TextField
                  name="age"
                  label="อายุ"
                  variant="outlined"
                  size="medium"
                  value={doctor.age || ''}
                  onChange={handleChangeNum}
                />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>โรคประจำตัว</InputLabel>
                <Select
                  name="disease"
                  label="โรคประจำตัว"
                  type="string"
                  value={doctor.disease || ''}
                  onChange={handleChange}
                >
                  {diseases.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.disease}</MenuItem>
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
                variant="outlined">
                <TextField
                  name="email"
                  label="อีเมล"
                  variant="outlined"
                  size="medium"
                  value={doctor.email || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  name="password"
                  label="รหัสผ่าน"
                  variant="outlined"
                  size="medium"
                  type={doctor.showPassword ? 'text' : 'password'}
                  value={doctor.password}
                  onChange={handleChange} />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  name="phone"
                  label="เบอร์ติดต่อ"
                  variant="outlined"
                  size="medium"
                  value={doctor.phone || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  name="address"
                  label="ที่อยู่ปัจจุบัน"
                  variant="outlined"
                  size="medium"
                  value={doctor.address || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  name="educational"
                  label="ประวัติการศึกษา"
                  variant="outlined"
                  size="medium"
                  value={doctor.educational || ''}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <div className={classes.formControl} style={{ marginLeft: 180 }}>
              <Button
                onClick={save}
                variant="contained"
                color="primary">
                บันทึกข้อมูล
             </Button>
              <Button style={{ marginLeft: 10 }}
                onClick={clear}
                variant="contained"
                color="secondary">
                CLEAR
             </Button>
              <Button style={{ marginLeft: 10 }}
                component={RouterLink} to="/table"
                variant="contained"
                color="secondary">
                BACK
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </div >
  );
};
export default Doctor;
