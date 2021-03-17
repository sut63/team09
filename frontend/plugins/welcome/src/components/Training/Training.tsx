import React, { FC, useEffect } from 'react';
import Button from '@material-ui/core/Button';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content
} from '@backstage/core';
import {
  AppBar,
  createStyles,
  FormControl,
  Grid,
  InputLabel,
  makeStyles,
  MenuItem,
  Select,
  TextField,
  Theme,
  Toolbar,
  Typography
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntCourse } from '../../api/models/EntCourse';
import { EntDoctor } from '../../api/models/EntDoctor';
import { EntDepartment } from '../../api/models/EntDepartment';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import Swal from 'sweetalert2'; // alert


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex', //ทำให้อยู่กึ่งกลาง
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    formControl: {
      margin: theme.spacing(1), //ขนาดของบล็อก
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

interface training {
  course: number;
  trainingplace: string;
  firstday: Date;
  lastday: Date;
  doctor: number;
  department: number;
  doctoridcard: string;
  hour: number;

}

const Training: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [training, setTraining] = React.useState<Partial<training>>({});
  const [courses, setCourses] = React.useState<EntCourse[]>([]); //setข้อมูล
  const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [trainingplaceError, setTrainingplaceError] = React.useState('');
  const [doctoridcardError, setDoctoridcardError] = React.useState('');
  const [hoursError, setHoursError] = React.useState('');


  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof Training;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setTraining({ ...training, [name]: value });
    console.log(training);
  };

  const handleChange1 = (
    event: React.ChangeEvent<{ name?: string; value: number }>,) => {
    const name = event.target.name as keyof typeof Training;
    const { value } = event.target;
    const validateValue = value.valueOf()
    checkPattern1(name, validateValue)
    setTraining({ ...training, [name]: +value });
    console.log(training);
  };

  const getCourses = async () => {   //ดึงข้อมูล
    const res = await http.listCourse({ limit: 10, offset: 0 });
    setCourses(res);
  };
  const getDoctors = async () => {
    const res = await http.listDoctor({ limit: 10, offset: 0 });
    setDoctors(res);
  };
  const getDepartments = async () => {
    const res = await http.listDepartment({ limit: 11, offset: 0 });
    setDepartments(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getCourses();
    getDoctors();
    getDepartments();
  }, []);

  // alert setting
  const Toast = Swal.mixin({
    position: 'center',
    showConfirmButton: true,
    timer: 5000
  });

  // clear input form
  function clear() {
    setTraining({});
  }

  const ValidateHours = (val: number) => {
    return val <= 100 && val >= 1 ? true : false;
  }

  const ValidateTrainingplace = (val: string) => {
    return val.match("^[ก-๏]+$")
  }
  const ValidateDoctoridcard = (val: string) => {
    return val.length == 10 ? true : false;
  }


  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'trainingplace':
        ValidateTrainingplace(value) ? setTrainingplaceError('') : setTrainingplaceError("กรอกสาขาให้เป็นภาษาไทย");
        return;

      case 'doctoridcard':
        ValidateDoctoridcard(value) ? setDoctoridcardError('') : setDoctoridcardError("กรอกตัวเลขทั้งหมด 10 ตัว");
        return;

      default:
        return;
    }
  }
  const checkPattern1 = (id: string, value: number) => {
    switch (id) {
      case 'hour':
        ValidateHours(Number(value)) ? setHoursError('') : setHoursError("กรอกชั่วโมง 1-100 เท่านั้น");
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
    switch (field) {
      case 'trainingplace':
        alertMessage("error", "กรอกสถานที่เข้าร่วมอบรมให้เป็นภาษาไทย");
        return;
      case 'doctoridcard':
        alertMessage("error", "กรอกตัวเลขทั้งหมด 10 ตัว");
        return;
      case 'hour':
        alertMessage("error", "กรอกชั่วโมง 1-100 เท่านั้น");
        return;
      default:
        alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ")
        return;


    }
  }


  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/trainings';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(training),
    };

    console.log(training); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
    <div className={classes.root} >
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" className={classes.title}>
            ระบบข้อมูลแพทย์
            </Typography>
          <Button color="inherit" href="/" startIcon={<ExitToAppRoundedIcon />}>Logout</Button>
        </Toolbar>
      </AppBar>

      <Content>
        <Typography variant="h4" className={classes.formControl} style={{ marginLeft: 100 }}>
          บันทึกข้อมูลการฝึกอบรม
      </Typography>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>หลักสูตรอบรม</InputLabel>
                <Select
                  name="course"
                  label="หลักสูตรอบรม"
                  type="string"
                  value={training.course || ''}
                  onChange={handleChange}
                >
                  {courses.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.namecourse}</MenuItem>
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
                  error={trainingplaceError ? true : false}
                  name="trainingplace"
                  label="สถานที่เข้าร่วมอบรม"
                  variant="outlined"
                  value={training.trainingplace || ''}
                  helperText={trainingplaceError}
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
                  label="วันแรกที่เข้าร่วมอบรม"
                  name="firstday"
                  type="date"
                  value={training.firstday || ''} // (undefined || '') = ''
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
                  label="วันแรกที่เข้าร่วมอบรม"
                  name="lastday"
                  type="date"
                  value={training.lastday || ''} // (undefined || '') = ''
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
                fullWidth
                className={classes.formControl}
                style={{ marginLeft: 100 }}
              >
                <TextField
                  error={hoursError ? true : false}
                  name="hour"
                  label="ชั่วโมงการอบรม"
                  variant="outlined"
                  value={training.hour || ''}
                  helperText={hoursError}
                  onChange={handleChange1}
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
                  error={doctoridcardError ? true : false}
                  name="doctoridcard"
                  label="เลขบัตรประจำตัวแพทย์"
                  variant="outlined"
                  value={training.doctoridcard || ''}
                  helperText={doctoridcardError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>


            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>ชื่อแพทย์</InputLabel>
                <Select
                  name="doctor"
                  label="ชื่อแพทย์"
                  type="string"
                  value={training.doctor || ''}
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
                  value={training.department || ''}
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
                variant="contained">
                BACK
             </Button>
            </div>
            </Grid>
            
          </form>
        </div>


      </Content>

    </div>
  );
};

export default Training;