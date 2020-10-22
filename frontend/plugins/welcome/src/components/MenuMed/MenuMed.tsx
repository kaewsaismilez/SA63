import React, { FC } from 'react';
import Button from '@material-ui/core/Button';
import ExitToAppRoundedIcon from '@material-ui/icons/ExitToAppRounded';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import Breadcrumbs from '@material-ui/core/Breadcrumbs';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Grid, { GridSpacing } from '@material-ui/core/Grid';



const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
    justifyContent: 'center'
  },
  paper: {
    height: 440,
    width: 1250,
    backgroundImage: 'url(https://image.freepik.com/free-vector/healthy-kawaii-tooth-with-dental-braces-cute-cartoon-character_74565-467.jpg)',
  },
  control: {
    padding: theme.spacing(2),
  },
 
})
);


const Homepage: FC<{}> = () => {
  const classes = useStyles();
  


  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ระบบบันทึกประวัติทันตกรรม">
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
      </Header>
      <Content>


      
        <ContentHeader title="Menu"> </ContentHeader>



        
        <Breadcrumbs aria-label="breadcrumb" >
            <Link 
            color="textPrimary" 
            href="/SaveMed" >
                บันทึกประวัติทันตกรรม
            </Link>
            <Link 
                color="textPrimary" 
                href="/ShowMed" 
            >
                แสดงข้อมูลประวัติทันตกรรม
            </Link>
        </Breadcrumbs>
        
        <br></br>
        <Grid container className={classes.root} spacing={2}>
      <Grid item xs={12}>
        <Grid container justify="center" >
          {[0].map((value) => (
            <Grid key={value} item>
              <Paper className={classes.paper} />
            </Grid>
          ))}
        </Grid>
      </Grid>
      <Grid item xs={12}>

      </Grid>
    </Grid>

        

       
        
      </Content>
    </Page>
  );
};
export default Homepage;
