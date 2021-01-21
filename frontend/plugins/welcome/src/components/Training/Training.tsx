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
  branch: string;
  dateone: Date;
  datetwo: Date;
  doctor: number;
  department: number;
  doctoridcard: string;
  hour: string;

}

const Training: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [training, setTraining] = React.useState<Partial<training>>({});
  const [courses, setCourses] = React.useState<EntCourse[]>([]); //setข้อมูล
  const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  const [branchError, setBranchError] = React.useState('');
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

  const getCourses = async () => {   //ดึงข้อมูล
    const res = await http.listCourse({ limit: 10, offset: 0 });
    setCourses(res);
  };
  const getDoctors = async () => {
    const res = await http.listDoctor({ limit: 10, offset: 0 });
    setDoctors(res);
  };
  const getDepartments = async () => {
    const res = await http.listDepartment({ limit: 10, offset: 0 });
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

  const ValidateHours = (val: string) => {
    return val.match("^[0-9]+$")
  }

  const ValidateBranch = (val: string) => {
    return val.match("^[ก-๏]+$")
  }
  const ValidateDoctoridcard = (val: string) => {
    return val.length == 10 ? true : false;
  }


  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'branch':
        ValidateBranch(value) ? setBranchError('') : setBranchError("กรอกสาขาให้เป็นภาษาไทย");
        return;

      case 'doctoridcard':
        ValidateDoctoridcard(value) ? setDoctoridcardError('') : setDoctoridcardError("กรอกตัวเลขทั้งหมด 10 ตัว");
        return;

      case 'hour':
        ValidateHours(value) ? setHoursError('') : setHoursError("กรอกชั่วโมงเป็นตัวเลขเท่านั้น");
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
    switch (field) {
      case 'branch':
        alertMessage("error", "กรอกสาขาให้เป็นภาษาไทย");
        return;
      case 'doctoridcard':
        alertMessage("error", "กรอกตัวเลขทั้งหมด 10 ตัว");
        return;
      case 'hour':
        alertMessage("error", "กรอกชั่วโมงเป็นตัวเลขเท่านั้น");
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
                  error={branchError ? true : false}
                  name="branch"
                  label="สาขา"
                  variant="outlined"
                  value={training.branch || ''}
                  helperText={branchError}
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
                  name="dateone"
                  type="date"
                  value={training.dateone || ''} // (undefined || '') = ''
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
                  label="วันสุดท้ายที่เข้าร่วมอบรม"
                  name="datetwo"
                  type="date"
                  value={training.datetwo || ''} // (undefined || '') = ''
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


            <div className={classes.formControl} style={{ marginLeft: 200 }}>
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
          </form>
        </div>
      </Content>

    </div>
  );
};

export default Training;