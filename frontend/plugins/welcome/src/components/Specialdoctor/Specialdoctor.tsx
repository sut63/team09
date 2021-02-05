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
  Extradoctor: number;
  Doctor: number;
  Department: number;
  Doctorid: string;
  Roomnumber: string;
  Other: string;
}

const Specialdoctor: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [doctors, setDoctor] = React.useState<EntDoctor[]>([]);
  const [departments, setDepartment] = React.useState<EntDepartment[]>([]);
  const [extradoctors, setExtradoctor] = React.useState<EntExtradoctor[]>([]);
  const [DoctoridError, setDoctoridError] = React.useState('');
  const [RoomnumberError, setRoomnumberError] = React.useState('');
  const [otherError, setOtherError] = React.useState('');

  const [specialdoctor, setSpecialdoctor] = React.useState<
    Partial<specialdoctor>>({});

  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 8000
  });
  //set data object path
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof Specialdoctor;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setSpecialdoctor({ ...specialdoctor, [name]: value });
    console.log(specialdoctor);
  };


  /*const handleChange2 = (
    event: React.ChangeEvent<{ name: string; value: number }>,
  ) => {
    const name = event.target.name as keyof typeof Specialdoctor;
    const { value } = event.target;
    setSpecialdoctor({ ...specialdoctor, [name]: +value });
  };*/

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

  }, []);

  function clear() {
    setSpecialdoctor({});
  }

  const ValidateOther = (val: string) => {
    return val.match("^[ก-๏\s]+$");
  }
  const ValidateRoomnumber= (val: string) => {
    const room = /^[ABC]\d{4}$/g
    return room.test(val);
  }
  const ValidateDoctorid = (val: string) => {
    return val.length == 10 ? true : false;
  }

  const checkPattern = (id: string, value: string) => {
    switch(id) {
      case 'Other' :
        ValidateOther(value) ? setOtherError('') : setOtherError("กรุณากรอกข้อมูลเฉพาะภาษาไทย");
        return;
        case 'Roomnumber' :
          ValidateRoomnumber(value) ? setRoomnumberError('') : setRoomnumberError("หมายเลขห้องขึ้นต้นด้วย A.B.C ตามด้วยตัวเลข 4 ตัว");
        return;
        case 'Doctorid' :
          ValidateDoctorid(value) ? setDoctoridError('') : setDoctoridError("กรอกตัวเลขทั้งหมด 10 ตัวห้ามกรอกตัวอักษร");
        return;
        default:
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
        case 'Other':
          alertMessage("error","กรุณากรอกข้อมูลเฉพาะภาษาไทย");
          return;
        case 'Roomnumber':
          alertMessage("error","หมายเลขห้องขึ้นต้นด้วย A.B.C ตามด้วยตัวเลข 4 ตัว");
          return;
        case 'Doctorid':
          alertMessage("error","กรอกตัวเลขทั้งหมด 10 ตัวห้ามกรอกตัวอักษร ");
          return;
        default:
          alertMessage("error","บันทึกข้อมูลไม่สำเร็จ")
          return;
        
    }
  }

  function save() {
    console.log(specialdoctor);
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
         checkCaseSaveError(data.error.Name)
          }
      });
  };

  return (
     <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบข้อมูลแพทย์
            </Typography>
            <Button color="inherit" href="/" startIcon={<ExitToAppRoundedIcon />}>Logout</Button>
        </Toolbar>
      </AppBar>
      <Content className={classes.withoutLabel}>
        <Typography variant="h5"
          className={classes.formControl} style={{ marginLeft: 150 }}>
          บันทึกข้อมูลแพทย์เฉพาะทาง
        </Typography>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>ชื่อ-นามสกุล</InputLabel>
                <Select
                  name="Doctor"
                  label="ชื่อ-นามสกุล"
                  type="string"

                  value={specialdoctor.Doctor || ''}
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
                  name="Department"
                  type="string"
                  label="แผนก"
                  value={specialdoctor.Department || ''}
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
                  name="Extradoctor"
                  label="สาขาเฉพาะทาง"
                  type="string"
                  value={specialdoctor.Extradoctor || ''}
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
                  error = {DoctoridError ? true : false}
                  name="Doctorid"
                  label="เลขบัตรประจำตัวแพทย์"
                  variant="outlined"
                  value={specialdoctor.Doctorid || ''}
                  helperText = {DoctoridError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
              >
                <TextField
                  error = {RoomnumberError ? true : false}
                  name="Roomnumber"
                  label="หมายเลขห้องทำงาน"
                  variant="outlined"
                  value={specialdoctor.Roomnumber || ''}
                  helperText = {RoomnumberError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
              >
                <TextField
                  error = {otherError ? true : false}
                  name="Other"
                  label="หมายเหตุ"
                  variant="outlined"
                  value={specialdoctor.Other || ''}
                  helperText = {otherError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid><br></br><br></br>
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