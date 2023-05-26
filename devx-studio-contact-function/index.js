const express = require('express');
const nodemailer = require('nodemailer');
const { MongoClient, ServerApiVersion } = require("mongodb");
require('dotenv').config();

const app = express();
const port = 8100;
const godaddyPassword = process.env.GODADDY_PASSWORD;
const godaddyEmail = process.env.GODADDY_EMAIL;
const godaddySMTP = process.env.GODADDY_SMTP;
const mongoURI = process.env.MONGO_URI;

const client = new MongoClient(mongoURI);

app.use(express.json());

app.post('/share-demo', (req, res) => {
    const { firstname,lastname ,email, company } = req.body;

    const transporter = nodemailer.createTransport({
    host: godaddySMTP,
    port: 465,
    secure: true,
    auth: {
      user: godaddyEmail,
      pass: godaddyPassword,
    },
  });

  const mailOptions = {
    from: godaddyEmail,
    to: email,
    subject: 'Test Email',
    text: `Hello ${firstname},\n\nThank you for your message: ${company}`,
    attachments: [
        {
            filename: 'feature-demo.pdf',
            path: './godaddy-studiodevx.pdf',
        },
        {
            filename: 'questions.pdf',
            path: './godaddy-studiodevx.pdf',
        }
    ]
  };

    transporter.sendMail(mailOptions, (error, info) => {
        if (error) {
            console.log(error);
            res.status(500).send('Something went wrong.');
        } else {
            console.log('Email sent: ' + info.response);
            res.status(200).send('Email successfully sent to recipient!');
        }
    });

    try {
        client.db("devx-contact").collection("potential-clients").insertOne({
            firstname: firstname,
            lastname: lastname,
            email: email,
            company: company,
        })
    } catch (error) {
        res.status(500).send('Something went wrong.');
    }
});

app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
  });