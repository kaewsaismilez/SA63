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

const Homepage: FC<{}> = () => {
  return (
    <Page theme={pageTheme.service}>
      <Header
        title={`Dental System`}
        subtitle="ระบบบันทึกประวัติทันตกรรม">
        <Button variant="contained" color="default" href="/" startIcon={<ExitToAppRoundedIcon />}> Logout
           </Button>
      </Header>
      <Content>
        <ContentHeader title="Menu">
        </ContentHeader>
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
      </Content>
    </Page>
  );
};
export default Homepage;