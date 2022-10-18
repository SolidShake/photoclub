import { Outlet } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';

import './App.css';

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
    primary: {
      light: '#f3e5f5',
      main: '#ce93d8',
      dark: '#ab47bc',
    },
    secondary: {
      light: '#e3f2fd',
      main: '#90caf9',
      dark: '#42a5f5',
    },
  },
});

export default function App() {
  return (
    <ThemeProvider theme={darkTheme}>
      <main className="main">
        <div className="main__container">
          <Outlet/>
        </div>
      </main>
    </ThemeProvider>
  );
}
