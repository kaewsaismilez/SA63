import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert
import { Link as RouterLink } from 'react-router-dom';

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Button,
  Link,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import { EntDentist } from '../../api/models/EntDentist'; // import interface Dentist
import { EntEmployee } from '../../api/models/EntEmployee'; // import interface Employee
import { EntPatient } from '../../api/models/EntPatient'; // import interface Patient

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface saveMed {
  Dentist: Number;
  Patient: Number;
  Employee: Number;
  Added: String;
  Detial: String;
}

const SaveMed: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [medicalfile, setMedicalfile] = React.useState<Partial<saveMed>>({});
  const [dentists, setDentists] = React.useState<EntDentist[]>([]);
  const [patients, setPatients] = React.useState<EntPatient[]>([]);
  const [employees, setEmployees] = React.useState<EntEmployee[]>([]);

  // alert setting
  //const Swal = require('sweetalert2')
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  // set data to object medicalfile
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof medicalfile;
    const { value } = event.target;
    setMedicalfile({ ...medicalfile, [name]: value });
    console.log(medicalfile);
  };
  

  const getPatient = async () => {
    const res = await http.listPatient({ limit: 4, offset: 0 });
    setPatients(res);
  };

  const getDentist = async () => {
    const res = await http.listDentist({ limit: 4, offset: 0 });
    setDentists(res);
  };

  const getEmployee = async () => {
    const res = await http.listEmployee({ limit: 4, offset: 0 });
    setEmployees(res);
  };

  // Lifecycle Hooks
  useEffect(() => {
    getPatient();
    getDentist();
    getEmployee();
  }, []);

  // clear input form
  function clear() {
    setMedicalfile({});
  }

  // function save data
  function save() {
    medicalfile.Added += ":00+07:00";
    const apiUrl = 'http://localhost:8080/api/v1/medicalfiles';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(medicalfile),
    };

    console.log(medicalfile); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {console.log(data.save);
        console.log(requestOptions)
        if (data.status == true) {
          clear();
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

  return (
    <Page theme={pageTheme.service}>
      <Header
       title="Dental System"
       subtitle="ระบบบันทึกประวัติทันตกรรม">
     </Header>
      <Content>
        <Container maxWidth="sm">
        {/* <div className='gameStatistics'>
            {
                Object.entries(medicalfile).map(([key, val]) => 
                    <h2 key={key}>{key}: {val}</h2>
                )
            }
        </div> */}
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className={classes.paper}>รหัสผู้ป่วย</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรหัสผู้ป่วย</InputLabel>
                <Select
                  name="Patient"
                  value={medicalfile.Patient || ''} 
                  onChange={handleChange}
                >
                  {patients.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.id}
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ทันตกรรม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <TextField 
                
                label="ประเภททันตกรรม" 
                variant="outlined" 
                name="Detial"
                type="string"
                value={medicalfile.Detial || ''}
                onChange={handleChange}
                />

              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ทันตแพทย์</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกทันตแพทย์ที่ทำการรักษา</InputLabel>
                <Select
                  name="Dentist"
                  value={medicalfile.Dentist || ''} 
                  onChange={handleChange}
                >
                  {dentists.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>วันที่ทำการรักษา</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  
                  label="เลือกวันที่"
                  name="Added"
                  type="datetime-local"
                  value={medicalfile.Added || ''} 
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>พนักงาน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกพนักงานที่ทำการบันทึก</InputLabel>
                <Select
                  name="Employee"
                  value={medicalfile.Employee || ''} 
                  onChange={handleChange}
                >
                  {employees.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>

                        {item.name}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกประวัติทันตกรรม
              </Button>
              &emsp;
              <Link component={RouterLink} to="/MenuMed">
              <Button
                variant="contained"
                color="default"
                size="large"
                style={{ marginLeft: 5 }}
              >
                กลับ
              </Button>
              </Link>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default  SaveMed;
