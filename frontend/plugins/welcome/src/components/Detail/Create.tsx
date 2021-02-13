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
import { EntMission } from '../../api/models/EntMission';
import { EntCourse } from '../../api/models/EntCourse';
import { EntDepartment } from '../../api/models/EntDepartment';
import Swal from 'sweetalert2';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
// import { KeyboardDatePicker, MuiPickersUtilsProvider } from '@material-ui/pickers';
// import DateFnsUtils from '@date-io/date-fns';

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
    textField: {
      marginLeft: 40,
      marginRight: theme.spacing(1),
      width: 200,
      minWidth: 300,
      minHeight: 100,

    },

  }),
);

interface detail {
  Department: number;
  Mission: number;
  Course: number;
  Explain: string;
  Phone: string;
  Email: string;
}

const Detail: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [detail, setDetail] = React.useState<Partial<detail>>({});
  const [departments, setDepartment] = React.useState<EntDepartment[]>([]);
  const [missions, setMission] = React.useState<EntMission[]>([]);
  const [courses, setCourse] = React.useState<EntCourse[]>([]);
  const [explainError, setexplainError] = React.useState('');
  const [emailError, setemailError] = React.useState('');
  const [phoneError, setphoneError] = React.useState('');

  const Toast = Swal.mixin({
    toast: true,
    position: 'center',
    showConfirmButton: false,
    timer: 2000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: any }>,) => {
    const name = event.target.name as keyof typeof Detail;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(name, validateValue)
    setDetail({ ...detail, [name]: value });
    // console.log(department);
  };

  const getDepartment = async () => {
    const res = await http.listDepartment({ limit: 11, offset: 0 });
    setDepartment(res);
  };
  const getCourse = async () => {
    const res = await http.listCourse({ limit: 10, offset: 0 });
    setCourse(res);
  };
  const getMission = async () => {
    const res = await http.listMission({ limit: 5, offset: 0 });
    setMission(res);
  };


  // Lifecycle Hooks
  useEffect(() => {
    getCourse();
    getMission();
    getDepartment();
    console.log(detail)
  }, [detail]);

  function clear() {
    setDetail({});
  }

  const Validateexplain = (val: string) => {
    return val.match("^[ก-๏\\s]+$");
  }
  const Validatephone = (val: string) => {
    return val.match("^[0-9]{10}$");
  }
  // const Validateemail = (val: string) => {
  //   return val.match("^[a-zA-Z_]+$");
  // }

  // ฟังก์ชั่นสำหรับ validate อีเมล
  const Validateemail = (email: string) => {
    const re = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(email);
  }

  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'Phone':
        Validatephone(value) ? setphoneError('') : setphoneError("กรุณากรอกเบอร์โทรศัพท์ 10 ตัวเท่านั้น");
        return;
      case 'Email':
        Validateemail(value) ? setemailError('') : setemailError("กรุณากรอกอีเมลให้ถูกต้อง");
        return;
      case 'Explain':
        Validateexplain(value) ? setexplainError('') : setexplainError("กรุณากรอกรายละเอียดเป็นภาษาไทยเท่านั้น");
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
      case 'phone':
        alertMessage("error", "กรุณากรอกเบอร์โทรศัพท์10ตัวเท่านั้น");
        return;
      
      case 'email':
        alertMessage("error", "กรุณากรอกอีเมลให้ถูกต้อง");
        return;
      
      case 'explain':
        alertMessage("error", "กรุณากรอกรายละเอียดเป็นภาษาไทยตัวเท่านั้น");
        return;
      
      default:
        alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ")
        return;
    }
  }

  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/details';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(detail),
    };

    console.log(detail); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
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
          <Button color="inherit" component={RouterLink} to="/" startIcon={<ExitToAppRoundedIcon />}> Logout </Button>
        </Toolbar>
      </AppBar>
      <Content className={classes.withoutLabel}>
        <Typography variant="h5"
          className={classes.formControl} style={{ marginLeft: 160 }}>
          บันทึกข้อมูลแผนกแพทย์
        </Typography>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>ชื่อแผนก</InputLabel>
                <Select
                  name="Department"
                  label="ชื่อแผนก"
                  type="string"
                  value={detail.Department || ''}
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
                <InputLabel>ตำแหน่ง/หน้าที่</InputLabel>
                <Select
                  name="Mission"
                  label="ตำแหน่ง/หน้าที่"
                  type="string"
                  value={detail.Mission || ''}
                  onChange={handleChange}
                >
                  {missions.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>{item.mission}</MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <InputLabel>หลักสูตร</InputLabel>
                <Select
                  name="Course"
                  label="หลักสูตร"
                  type="string"
                  value={detail.Course || ''}
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
                variant="outlined">
                <TextField
                  error={phoneError ? true : false}
                  name="Phone"
                  label="เบอร์ติดต่อ"
                  variant="outlined"
                  size="medium"
                  value={detail.Phone || ''}
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
                  error={emailError ? true : false}
                  name="Email"
                  label="อีเมล"
                  variant="outlined"
                  size="medium"
                  value={detail.Email || ''}
                  helperText={emailError}
                  onChange={handleChange}
                />
              </FormControl>
            </Grid>

            <Grid item xs={6}>
              <FormControl variant="outlined" className={classes.formControl} style={{ marginLeft: 100 }}>
                <TextField
                  error={explainError ? true : false}
                  name="Explain"
                  value={detail.Explain || ''}
                  label="ข้อมูลแผนก"
                  multiline
                  rows={5}
                  variant="outlined"
                  helperText={explainError}
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
export default Detail;