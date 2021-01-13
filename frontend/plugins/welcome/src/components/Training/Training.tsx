import React, { FC, useEffect } from 'react';
import Button from '@material-ui/core/Button'; 
import {
 Content
} from '@backstage/core';
import { AppBar, 
    createStyles, 
    FormControl, 
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
      minWidth: 350,
      maxWidth: 300,
      marginTop: '-1%',
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    title: {
      flexGrow: 1,
    },
    title2: {
      flexGrow: 1,
      marginLeft:30,
    },
   
    textField: {
      marginLeft: theme.spacing(1),
      marginRight: theme.spacing(1),
      width: 330,
    },
    textField2: {
      marginLeft: theme.spacing(0),
      marginRight: theme.spacing(1),
      width: 348,}
  }),
);
 
interface training {
    course: number;
    branch: string;
    dateone: Date;
    datetwo: Date;
    doctor:  number;
    department: number;

}

const Training: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();
  const [training, setTraining] = React.useState<Partial<training>>({});
  const [courses, setCourses] = React.useState<EntCourse[]>([]); //setข้อมูล
  const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
  const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
  
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
        const name = event.target.name as keyof typeof Training;
        const { value } = event.target;
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
          Toast.fire({
            title: 'บันทึกข้อมูลไม่สำเร็จ',
            icon: 'error',
          });
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
      <Typography variant="h4" className={classes.title2}>
          บันทึกข้อมูลการฝึกอบรม
      </Typography>
      <br></br>
 
      <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >หลักสูตรอบรม</InputLabel>
        <Select
          name="course"
          type="string"     
          value={training.course || ''} 
          onChange={handleChange}
          label="หลักสูตรอบรม"
          >
          {courses.map(item => {
            return(<MenuItem key={item.id} value={item.id}>{item.namecourse}</MenuItem>
                );
            })}</Select>
      </FormControl> <br></br><br></br>

      <form className={classes.formControl} noValidate >
      <TextField 
      name="branch"
      type="string" 
      label="สาขา" 
      variant="outlined" 
      className={classes.textField2}
      value={training.branch || '' }
      onChange={handleChange}
      />
      </form><br></br><br></br>

      <form className={classes.formControl} noValidate>
      <TextField
        name="dateone"
        label="วันแรกที่เข้าร่วมอบรม"
        type="date"
        className={classes.textField}
        value={training.dateone || '' } 
        InputLabelProps={{
          shrink: true,
        }}
        onChange={handleChange}
      />
    </form><br></br><br></br>

    <form className={classes.formControl} noValidate>
      <TextField
        name="datetwo"
        label="วันสุดท้ายที่เข้าร่วมอบรม"
        type="date"
        className={classes.textField}
        value={training.datetwo  || ''} 
        InputLabelProps={{
          shrink: true,
        }}
        onChange={handleChange}
      />
    </form><br></br><br></br>

      <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >ชื่อ - นามสกุล</InputLabel>
        <Select
          name="doctor"
          type="string"
          onChange={handleChange}
          label="ชื่อ - นามสกุล"
          value={training.doctor || ''} 
        >
        {doctors.map(item => {
         return (
            <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                );
            })}

        </Select>
      </FormControl> <br></br><br></br>

      <FormControl variant="outlined" className={classes.formControl}>
        <InputLabel >แผนก</InputLabel>
        <Select
          name="department"
          type= "string"
          onChange={handleChange}
          value={training.department || ''}
          label="แผนก">
          {departments.map(item => {
         return (
            <MenuItem key={item.id} value={item.id}>{item.name}</MenuItem>
                );
            })}
        </Select>
      </FormControl> <br></br><br></br>

      <Button variant="contained" color="primary"  onClick={save}  style={{ marginRight: 70 }}>SAVE</Button>
      <Button variant="contained" color="secondary" onClick={clear}  style={{ marginRight: 70 }}>CLEAR</Button>
      <Button variant="contained" href="/home" >BACK</Button>
      </Content>

   </div>
 );
};
 
export default Training;