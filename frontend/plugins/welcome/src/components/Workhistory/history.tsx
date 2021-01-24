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

interface office {
  roomnumber: string;
  doctoridcard: string;
  officename: string;
  doctor: number;
  department: number;
  specialdoctor: number;
  extradoctor: number;
  added1: Date;
  added2: Date;
}

const Office: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [office, setOffice] = React.useState<Partial<office>>({});
  const [doctors, setDoctor] = React.useState<EntDoctor[]>([]);
  const [officenameError, setOfficenameError] = React.useState('');
  const [roomnumberError, setRoomnumberError] = React.useState('');
  const [doctoridcardError, setDoctoridcardError] = React.useState('');
  const [departments, setDepartment] = React.useState<EntDepartment[]>([]);
  const [extradoctors, setExtradoctor] = React.useState<EntExtradoctor[]>([]);

  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 9000
  });

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof Office;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setOffice({ ...office, [name]: value });
    console.log(office);
  };

  const getDoctors = async () => {
    const res = await http.listDoctor({ limit: 10, offset: 0 });
    setDoctor(res);
  };
  const getDepartments = async () => {
    const res = await http.listDepartment({ limit: 10, offset: 0 });
    setDepartment(res);
  };
  const getExtradoctors = async () => {
    const res = await http.listExtradoctor({ limit: 10, offset: 0 });
    setExtradoctor(res);
  };
  
  // Lifecycle Hooks
  useEffect(() => {
    getDoctors();
    getDepartments();
    getExtradoctors();
  }, []);

  function clear() {
    setOffice({});
  }
  const ValidateOfficename = (val: string) => {
    return val.match("^[ก-๏\s]+$");
  }
  const ValidateRoomnumber = (val: string) => {
    return val.match("[ABC]\\d{4}");
  }
  const ValidateDoctoridcard = (val: string) => {
    return val.length == 10 ? true : false;
  }

  const checkPattern = (id: string, value: string) => {
    switch(id) {
      case 'officename' :
        ValidateOfficename(value) ? setOfficenameError('') : setOfficenameError("กรอกสถานที่ทำงานให้เป็นภาษาไทย");
        return;
        case 'roomnumber' :
          ValidateRoomnumber(value) ? setRoomnumberError('') : setRoomnumberError("หมายเลขห้องขึ้นต้นด้วย A.B.C ตามด้วยตัวเลข 4 ตัว");
        return;
        case 'doctoridcard' :
          ValidateDoctoridcard(value) ? setDoctoridcardError('') : setDoctoridcardError("กรอกตัวเลขทั้งหมด 10 ตัวห้ามกรอกตัวอักษร");
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
        case 'officename':
          alertMessage("error","กรอกสถานที่ทำงานให้เป็นภาษาไทยเท่านั้น");
          return;
        case 'roomnumber':
          alertMessage("error","หมายเลขห้องขึ้นต้นด้วย A.B.C ตามด้วยตัวเลข 4 ตัว");
          return;
        case 'doctoridcard':
          alertMessage("error","กรอกตัวเลขทั้งหมด 10 ตัว");
          return;
        default:
          alertMessage("error","บันทึกข้อมูลไม่สำเร็จ")
          return;
        
    }
  }
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/offices';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(office),
    };
    console.log(office); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
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
  }
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
          className={classes.formControl} style={{ marginLeft: 100 }}>
          บันทึกประวัติการทำงานของแพทย์
        </Typography>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>ชื่อแพทย์</InputLabel>
                <Select
                  name="doctor"
                  label="ชื่อแพทย์"
                  type="string"
                  value={office.doctor || ''}
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
                  label="แผนก"
                  type="string"
                  value={office.department || ''}
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
                <InputLabel>แพทย์เฉพาะทาง</InputLabel>
                <Select
                  name="extradoctor"
                  label="แพทย์เฉพาะทาง"
                  type="string"
                  value={office.extradoctor || ''}
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
                  error = {officenameError ? true : false}
                  name="officename"
                  label="สถานที่ทำงาน"
                  variant="outlined"
                  value={office.officename || ''}
                  helperText = {officenameError}
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
                  error = {doctoridcardError ? true : false}
                  name="doctoridcard"
                  label="เลขบัตรประจำตัวแพทย์"
                  variant="outlined"
                  value={office.doctoridcard || ''}
                  helperText = {doctoridcardError}
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
                  error = {roomnumberError ? true : false}
                  name="roomnumber"
                  label="หมายเลขห้องทำงาน"
                  variant="outlined"
                  value={office.roomnumber|| ''}
                  helperText = {roomnumberError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl
                style={{ marginLeft: 95 }}
                variant="outlined"
                className={classes.formControl}
              >
                <TextField
                  label="ระยะเวลาในการทำงาน - วันแรก"
                  name="added1"
                  type="date"
                  value={office.added1 || ''} // (undefined || '') = ''
                  className={classes.formControl}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <Grid item xs={6}>
              <FormControl
                style={{ marginLeft: 95 }}
                variant="outlined"
                className={classes.formControl}
              >
                <TextField
                  label="ระยะเวลาในการทำงาน - ปัจจุบัน"
                  name="added2"
                  type="date"
                  value={office.added2 || ''} // (undefined || '') = ''
                  className={classes.formControl}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>
            <div className={classes.formControl} style={{ marginLeft: 110 }}>
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
             <Button style={{ marginLeft: 10 }}
                component={RouterLink} to="/tables"
                variant="contained"
                color="secondary">
                SHOWTABLE
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </div>
  );
};
export default Office;