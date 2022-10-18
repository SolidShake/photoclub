import { render, screen } from '@testing-library/react';
import App from './App';

test('renders main', () => {
  render(<App />);
  const appElement = screen.getByTestId('app');
  expect(appElement).toBeInTheDocument();
});
