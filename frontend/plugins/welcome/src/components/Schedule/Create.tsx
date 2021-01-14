// import { Link as RouterLink } from 'react-router-dom';
import {
    Content,
    // ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import TextField from '@material-ui/core/TextField'
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Select from '@material-ui/core/Select';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Grid from '@material-ui/core/Grid';
import React, { useEffect, FC } from 'react';
import { EntDepartment } from '../../api/models/EntDepartment';
import { EntOffice } from '../../api/models/EntOffice';
import { EntDoctor } from '../../api/models/EntDoctor';
import { DefaultApi } from '../../api/apis';
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
            minWidth: 300,
            minHeight: 100,
            marginTop: '5%',
            marginLeft: 40,

        },
        container: {
            display: 'flex',
            flexWrap: 'wrap',
            minWidth: 300,


        },
        textField: {
            marginLeft: 40,
            marginRight: theme.spacing(1),
            width: 200,
            minWidth: 300,
            minHeight: 100,

        },
        menuButton: {

            marginLeft: 500,
        },
        title: {
            flexGrow: 1,
            marginTop: '5%',
        },
        Bottom: {
            minWidth: 300,
            minHeight: 100,
        }

    }),
);

interface schedule {
    Doctor: number;
    Department: number;
    Office: number;
    Activity: number;
    Added: Date;
}

const Schedule: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    const [schedule, setSchedule] = React.useState<Partial<schedule>>({});
    const [departments, setDepartments] = React.useState<EntDepartment[]>([]);
    const [offices, setOffices] = React.useState<EntOffice[]>([]);
    const [doctors, setDoctors] = React.useState<EntDoctor[]>([]);
    // const [alert, setAlert] = React.useState(true);


    // alert setting
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


    // const handleAmountChange = (
    //   event: React.ChangeEvent<{ name: string; value: number }>,) => {
    //   const name = event.target.name as keyof typeof Schedule;
    //   const { value } = event.target;
    //   setSchedule({ ...schedule, [name]: +value });
    //   // console.log(schedule);
    // };
    // const handleDateChange = (
    //     event: React.ChangeEvent<{ name: string; value: string }>,) => {
    //     const name = event.target.name as keyof typeof Schedule;
    //     const { value } = event.target;
    //     console.log('date select: ', value, typeof value) // show date from event.target.value
    //     setSchedule({ ...schedule, [name]: value });
    //     // console.log(schedule);
    // };
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: unknown }>,) => {
        // console.log(schedule);
        const name = event.target.name as keyof typeof Schedule;
        const { value } = event.target;
        setSchedule({ ...schedule, [name]: value });
        // console.log(schedule);
    };


    const getDepartments = async () => {
        const res = await api.listDepartment({ limit: 10, offset: 0 });
        setDepartments(res);
    };

    const getOffices = async () => {
        const res = await api.listOffice({ limit: 10, offset: 0 });
        setOffices(res);
    };

    const getDoctors = async () => {
        const res = await api.listDoctor({ limit: 10, offset: 0 });
        setDoctors(res);
    };

    // Lifecycle Hooks
    useEffect(() => {
        getDepartments();
        getOffices();
        getDoctors();
        console.log(schedule)
    }, [schedule]);

    // clear input form
    function clear() {
        setSchedule({});
    }

    function save() {
        schedule.Added + ":00+07:00";
        debugger
        const apiUrl = 'http://localhost:8080/api/v1/schedules';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(schedule),
        };
        // console.log('date: ', schedule.Added)
        console.log('schedule: ', schedule)
        // console.log(schedule); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                // setSchedule({ ...schedule, Added: schedule.Added });
                // console.log(added_time)
                console.log(data); // last data
                // alert(JSON.stringify(schedule))
                if (data.status === true) {
                    Toast.fire({
                        icon: 'success',
                        title: 'บันทึกข้อมูลสำเร็จ',
                    });
                } else {
                    Toast.fire({
                        icon: 'error',
                        title: 'บันทึกข้อมูลไม่สำเร็จ',
                    });
                }
            });
    }

    function LogOut() {
        //redirec Page ... http://localhost:3000/
        window.location.href = "http://localhost:3000/";
    }


    return (
        <div className={classes.root}>

            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" className={classes.root}>
                        ระบบบันทึกเวลาของแพทย์
              </Typography>
                    <Grid item xs> </Grid>
                    <Button color="inherit" href="/" startIcon={<ExitToAppRoundedIcon />}>Logout</Button>
                </Toolbar>
            </AppBar>

            <Content>
                <Typography variant="h3" className={classes.title}>
                    ระบบบันทึกเวลาของแพทย์
              </Typography>

                {/* ชื่อแพทย์ */}
                <div className={classes.formControl}>
                    <form noValidate autoComplete="off">
                        <Grid item xs>
                            <FormControl variant="outlined" className={classes.formControl}>
                                <InputLabel>ชื่อแพทย์</InputLabel>
                                <Select
                                    name="Doctor"
                                    label="ชื่อแพทย์"
                                    value={schedule.Doctor || ''}
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

                        {/* แผนก */}
                        <Grid item xs>

                            <FormControl variant="outlined" className={classes.formControl} >
                                <InputLabel>แผนก</InputLabel>
                                <Select
                                    name="Department"
                                    label="คลังยา"
                                    value={schedule.Department || ''}
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


                        {/* กิจกกรม */}
                        <Grid item xs>
                            <FormControl variant="outlined" className={classes.textField} >
                                {/* <form className={classes.textField} noValidate autoComplete="off"> */}
                                <TextField
                                    name="Activity"
                                    label="กิจกรรมของแพทย์"
                                    id="Activity"
                                    type="string"
                                    value={schedule.Activity || ''}

                                    onChange={handleChange}

                                />
                                {/* </form> */}
                            </FormControl>
                        </Grid>

                        {/* วันที่ */}
                        <Grid item xs>

                            {/* <form className={classes.container} noValidate> */}
                            <TextField
                                name="Added"
                                label="วันที่-เวลา"
                                type="datetime-local"
                                value={schedule.Added || ''} // (undefined || '') = ''
                                className={classes.textField}
                                InputLabelProps={{
                                    shrink: true,
                                }}
                                onChange={handleChange}
                            />
                            {/* </form> */}
                        </Grid>

                        <Grid item xs>
                            {/* สถานที่ทำงาน */}
                            <FormControl variant="outlined" className={classes.formControl}>
                                <InputLabel>สถานที่ทำงาน</InputLabel>
                                <Select
                                    name="Office"
                                    label="สถานที่ทำงาน"
                                    value={schedule.Office}
                                    onChange={handleChange}
                                >
                                    {offices.map(item => {
                                        return (
                                            <MenuItem key={item.id} value={item.id}>{item.officename}</MenuItem>
                                        );
                                    })}
                                </Select>
                            </FormControl>
                        </Grid>

                        {/* ปุ่ม */}

                        <Button variant="contained" color="primary" onClick={save} style={{ marginLeft: 100 }}>บันทึกข้อมูล</Button>
                        <Button variant="contained" color="secondary" onClick={clear} style={{ marginLeft: 10 }}>ยกเลิก</Button>

                    </form>
                </div>
            </Content>
        </div>

    );
};

export default Schedule;
