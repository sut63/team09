import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title = {'ยินดีตอนรับเข้าสู้ระบบห้องยา'}
      //  subtitle=""
     ></Header>
     <Content>
       <ContentHeader title="ข้อมูลเภสัชกร">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             บันทึกข้อมูลเภสัชกร
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;