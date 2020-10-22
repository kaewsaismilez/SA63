import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import ComponanceTable from '../Table';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';


const useStyles = makeStyles((theme: Theme) =>
 createStyles({
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(70),
    },
  }),
);

const ShowMed: FC<{}> = () => {
  //const classes = useStyles();
 return (
   
   <Page theme={pageTheme.service}>
     <Header
       title="Dental System"
       subtitle="ระบบบันทึกประวัติทันตกรรม">
     </Header>

      
        <Content>
        <ContentHeader title="ประวัติทันตกรรม">
        <Link component={RouterLink} to="/MenuMed">
              <Button
                variant="contained"
                color="default"
              >
                กลับ
              </Button>
              </Link>
              &emsp;
        <Link component={RouterLink} to="/SaveMed">
            <Button variant="contained" color="primary">
              บันทึกข้อมูลทันตกรรมผู้ป่วย
            </Button>
          </Link>
        </ContentHeader>
        <ComponanceTable></ComponanceTable>
        </Content>
      
   </Page>
 );
};

export default ShowMed;