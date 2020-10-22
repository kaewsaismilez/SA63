import { createPlugin } from '@backstage/core';
import SignIn from './components/SignIn';
import ShowMed from './components/ShowMed';
import MedicalfileCreate from './components/SaveMed';
import MenuMed from './components/MenuMed';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
    router.registerRoute('/MenuMed', MenuMed);
    router.registerRoute('/ShowMed', ShowMed);
    router.registerRoute('/SaveMed', MedicalfileCreate);
    
  }
});
