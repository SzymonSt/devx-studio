import React, { useEffect, useState } from 'react';
import logo from './logo.svg';
import Form from 'react-bootstrap/Form';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';
import './DemoRequest.css';

function DemoRequest() {
  const $apiUrl = process.env.REACT_APP_API_URL;

  const [name, setName] = useState('');
  const [lastname, setLastname] = useState('');
  const [email, setEmail] = useState('');
  const [company, setCompany] = useState('');
  const [validated, setValidated] = useState(false);
  const [buttonClassname, setButtonClassName ] = useState('demo-button-disabled');


  const handleNameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setName(e.target.value);
  }

  const handleLastnameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLastname(e.target.value);
  }

  const handleEmailChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setEmail(e.target.value);
  }

  const handleCompanyChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setCompany(e.target.value);
  }

  useEffect(() => {
    if (name && lastname && email && company) {
      setValidated(true);
      setButtonClassName('demo-button');
    }else{
      setValidated(false);
      setButtonClassName('demo-button-disabled');
    }
  }, [name, lastname, email, company]);

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    async function submitForm(){
    await fetch( $apiUrl + '/share-demo', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            "firstname": name,
            "lastname": lastname,
            "email": email,
            "company":company
          })
        }
      ).then((response) => {
        if (response.status === 200) {
          alert('Thank you for your interest! We will be in touch shortly.');
          setName('');
          setLastname('');
          setEmail('');
          setCompany('');
        }else{
          alert('Something went wrong. Please try again later.');
        }
      }).catch((error) => {
        alert('Something went wrong. Please try again later.');
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
            <p className='demo-intro'>📢 Attention! Have you been searching tirelessly for a tool that can truly engage developers and measure the impact of your actions on their experience? Are you struggling to determine whether your efforts are making a difference in how developers work and utilize your platform?</p>
            <p className='demo-intro'>By leaving your contact details below, you'll gain access to:
             <ul>
              <li><span className='emph'>Short demo that showcases how DevXStudio features can revolutionize your approach to developer experience surveys.</span></li>
              <li><span className='emph'>Comprehensive guide featuring 10 essential questions you can ask your developers to assess the state of DevEx within your organization.</span></li>
             </ul>
            </p>
            <p className='demo-intro'>🚀 Get ready to supercharge developer experience within your organization!</p>
          </span>
          <Form.Group className='form-gr' controlId="formBasicEmail">
            <Form.Label className='form-label'>Firstname<span className='astrix'>*</span></Form.Label>
            <Form.Control value={name} onChange={handleNameChange} className='demo-input' type="text" placeholder="Enter name" />
            <Form.Label className='form-label'>Lastname<span className='astrix'>*</span></Form.Label>
            <Form.Control value={lastname} onChange={handleLastnameChange} className='demo-input' type="text" placeholder="Enter last name" />
            <Form.Label className='form-label'>Email<span className='astrix'>*</span></Form.Label>
            <Form.Control value={email} onChange={handleEmailChange} className='demo-input' type="email" placeholder="Enter email" />
            <Form.Label className='form-label'>Company<span className='astrix'>*</span></Form.Label>
            <Form.Control value={company} onChange={handleCompanyChange} className='demo-input' type="text" placeholder="Enter company name" />
            <Button disabled={!validated} className={buttonClassname} variant="primary" type="submit">Submit</Button>
          </Form.Group>
        </Form>
      </div>
    </div>
  );
}

export default DemoRequest;
