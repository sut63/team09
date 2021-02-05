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
// import { EntDoctor } from '../../api/models/EntDoctor';

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
  const [nameError, setnameError] = React.useState('');
  const [passwordError, setpasswordError] = React.useState('');
  const [phoneError, setphoneError] = React.useState('');
  const [addressError, setaddressError] = React.useState('');
  const [educationalError, seteducationalError] = React.useState('');
  const [emailError, setemailError] = React.useState('');
  const [ageError, setageError] = React.useState('');
  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 1500
  });

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof Doctor;
    const { value } = event.target;
    setDoctor({ ...doctor, [name]: value });
    const validateValue = value.toString()
    checkPattern(name, validateValue)

    console.log(doctor);
  };

  const handleChangeNum = (
    event: React.ChangeEvent<{ name: string; value: number }>,) => {
    const name = event.target.name as keyof typeof Doctor;
    const { value } = event.target;
    const validateValue = value.valueOf()
    checkPattern1(name, validateValue)
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
  
  const Validateemail = (email: string) => {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
  }
  const Validatename = (val: string) => {
    return val.match("^[ก-๏\\s]+$");
  }
  const Validatephone = (val: string) => {
    return val.match("^[0-9]{10}$");
  }
  const Validatepassword = (val: string) => {
    return val.match("^[0-9]{8}$");
  }
  const Validateaddress = (val: string) => {
    return val.match("");
  }
  const Validateeducational = (val: string) => {
    return val.match("");
  }
  const Validateage = (val: number) => {
    return val <= 55 && val >= 25 ? true : false;
  }


  const checkPattern = (id: string, value: string) => {
    switch(id) {
      case 'phone' :
        Validatephone(value) ? setphoneError('') : setphoneError("กรุณากรอกเบอร์โทรศัพท์10ตัวเท่านั้น");
        return;  
      case 'name' :    
        Validatename(value) ? setnameError('') : setnameError("กรุณากรอกชื่อแพทย์เป็นภาษาไทยเท่านั้น");
        return;
      case 'password' :
        Validatepassword(value) ? setpasswordError('') : setpasswordError("กรุณากรอกรหัสผ่านตัวเลข0-9จำนวน8ตัวเท่านั้น");
        return;
      case 'address' :
          Validateaddress(value) ? setaddressError('') : setaddressError("กรุณากรอกที่อยู่");
          return;
      case 'educational' :
            Validateeducational(value) ? seteducationalError('') : seteducationalError("กรุณากรอกประวัติการศึกษา");
            return;
      case 'email':
            Validateemail(value) ? setemailError('') : setemailError("กรุณากรอกอีเมลให้ถูกต้องตามรูปแบบ xxx@x.xxx");
            return;
      // case 'age':
      //       Validateage(value) ? setageError('') : setageError("กรุณาอายุเป็นจำนวนเต็มบวกเท่านั้น");
      //       return;
        default:
          return;
    }
  }
  const checkPattern1 = (id: string, value: number) => {
    switch (id) {
      case 'age':
        Validateage(Number(value)) ? setageError('') : setageError("กรอกอายุตั้งแต่ 26-55 ปีเท่านั้น");
        return;
    }
  }

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
        case 'phone':
          alertMessage("error","กรุณากรอกเบอร์โทรศัพท์10ตัวเท่านั้น");
          return;
        case 'name':
            alertMessage("error","กรุณากรอกชื่อแพทย์เป็นภาษาไทยเท่านั้น");
            return;
        case 'password':
              alertMessage("error","กรุณากรอกรหัสผ่านตัวเลข0-9จำนวน8ตัวเท่านั้น");
              return;
        case 'address':
              alertMessage("error","กรุณากรอกที่อยู่");
              return;
        case 'educational':
              alertMessage("error","กรุณากรอกประวัติการศึกษา");
              return;
        case 'email':
              alertMessage("error","กรุณากรอกอีเมลให้ถูกต้อง");
              return;
         case 'age':
              alertMessage("error","กรอกอายุตั้งแต่ 26-55 ปีเท่านั้น");
              return;
         default:
              alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ")
              return;
        
    }
  }

  
  function save() {
    if(doctor.age){
      var age: number = +doctor.age;
      doctor.age = age;
    }
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
            checkCaseSaveError(data.error.Name)
        }
          });
        
      };
  

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
          className={classes.formControl} style={{ marginLeft: 350 }}>
          บันทึกข้อมูลส่วนตัวแพทย์
        </Typography>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <Grid container spacing={3}>
            <Grid item xs={6}>
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
              <Grid item xs={6}>
                <FormControl
                  fullWidth
                  className={classes.formControl}
                  style={{ marginLeft: 100 }}
                  variant="outlined">
                  <TextField
                    error = {nameError ? true : false}
                    name="name"
                    label="ชื่อแพทย์"
                    variant="outlined"
                    size="medium"
                    value={doctor.name || ''}
                    helperText = {nameError}
                    onChange={handleChange}
                  />
                </FormControl>
              </Grid>  
              </Grid>   
              <Grid container spacing={3}>
              <Grid item xs={6}>
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
            </Grid>
            <Grid container spacing={3}>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  error={ageError ? true : false}
                  name="age"
                  label="อายุ"
                  variant="outlined"
                  size="medium"
                  value={doctor.age || ''}
                  helperText={ageError}
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
            </Grid>
            <Grid container spacing={3}>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  error={emailError ? true : false}
                  name="email"
                  label="อีเมล"
                  variant="outlined"
                  size="medium"
                  value={doctor.email || ''}
                  helperText={emailError}
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
                  error={passwordError ? true : false}
                  name="password"
                  label="รหัสผ่าน"
                  variant="outlined"
                  size="medium"
                  type={doctor.showPassword ? 'text' : 'password'}
                  value={doctor.password || ''}
                  helperText={passwordError}
                  onChange={handleChange} />
              </FormControl>
            </Grid>
            </Grid>
            <Grid container spacing={3}>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                  error={phoneError ? true : false}
                  name="phone"
                  label="เบอร์ติดต่อ"                  
                  variant="outlined"
                  size="medium"
                  value={doctor.phone || ''}
                  helperText={phoneError}
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
                  error={addressError ? true : false}
                  name="address"
                  label="ที่อยู่ปัจจุบัน"
                  variant="outlined"
                  size="medium"
                  value={doctor.address || ''}
                  helperText={addressError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            </Grid>
            <Grid container spacing={3}>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
                variant="outlined">
                <TextField
                 error={educationalError ? true : false}
                  name="educational"
                  label="ประวัติการศึกษา"
                  variant="outlined"
                  size="medium"
                  value={doctor.educational || ''}
                  helperText={educationalError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            </Grid>
            <div className={classes.formControl} style={{ marginLeft: 450 }}>
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
            </div>
          </form>
        </div>
      </Content>
    </div >
  );
};
export default Doctor;
