import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import { useNavigate } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Waitlist.css';

function Waitlist() {
  const $apiUrl = process.env.REACT_APP_API_URL;

  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [validated, setValidated] = useState(false);
  const [buttonClassname, setButtonClassName ] = useState('demo-button-disabled');
  const navigate = useNavigate();

  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  }

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  }

  useEffect(() => {
    if (name && email) {
      setValidated(true);
      setButtonClassName('demo-button');
    }else{
      setValidated(false);
      setButtonClassName('demo-button-disabled');
    }
  }, [name, email]);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    async function submitForm(){
    await fetch( $apiUrl + '/waitlist', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            "firstname": name,
            "email": email
          })
        }
      ).then((response) => {
        if (response.status === 200) {
          setName('');
          setEmail('');
          navigate('/thanks-for-interest');
        }else{
          navigate('/something-went-wrong');
        }
      }).catch((error) => {
        navigate('/something-went-wrong');
      }
      );
    }
    submitForm();
  }
  

  return (
    <div className="App">
      <div className="App-header">
        <h2 className="logo">
          <span className='logo-p1'>DevX</span>
          <span className='logo-p2'>Studio</span>
        </h2>
      </div>
      <div className="Form-box">
        <Form onSubmit={handleSubmit}>
          <span>
            <p className='demo-intro'>Sign up for <span className='logo-p1'>DevX</span><span className='logo-p2'>Studio</span> waitlist. We will keep you posted about devxstudio release and features.</p>
          </span>
          <Form.Group className='form-gr' controlId="formBasicEmail">
            <Form.Label className='form-label'>Firstname<span className='astrix'>*</span></Form.Label>
            <Form.Control value={name} onChange={handleNameChange} className='demo-input' type="text" placeholder="Enter name" />
            <Form.Label className='form-label'>Email<span className='astrix'>*</span></Form.Label>
            <Form.Control value={email} onChange={handleEmailChange} className='demo-input' type="email" placeholder="Enter email" />
            <Button disabled={!validated} className={buttonClassname} variant="primary" type="submit">Submit</Button>
          </Form.Group>
        </Form>
      </div>
    </div>
  );
}

export default Waitlist;
