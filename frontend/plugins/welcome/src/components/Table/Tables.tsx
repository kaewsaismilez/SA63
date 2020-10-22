import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
//import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntMedicalfile } from '../../api/models';
import moment from 'moment';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});

export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [medicalfile, setMedicalfile] = useState<EntMedicalfile[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const getMedicalfile = async () => {
      const res = await api.listMedicalfile({ limit: 10, offset: 0 });
      setLoading(false);
      setMedicalfile(res);
      console.log(res);
    };
    getMedicalfile();
  }, [loading]);

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell align="center">Medicalfile_ID</TableCell>
            <TableCell align="center">Patient</TableCell>
            <TableCell align="center">Detail</TableCell>
            <TableCell align="center">Dentist</TableCell>
            <TableCell align="center">Date</TableCell>
            <TableCell align="center">Employee</TableCell>
          </TableRow>
        </TableHead>
        
        <TableBody>
          {medicalfile.map((item:any) => (
            <TableRow key={item.id}>
              <TableCell align="center">{item.id}</TableCell>
              <TableCell align="center">{item.edges?.patient?.name}</TableCell>
              <TableCell align="center">{item.detail}</TableCell>
              <TableCell align="center">{item.edges?.dentist?.name}</TableCell>
              <TableCell align="center">{moment(item.addedTime).format('DD/MM/YYYY HH:mm')}</TableCell>
              <TableCell align="center">{item.edges?.employee?.name}</TableCell>
              
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
