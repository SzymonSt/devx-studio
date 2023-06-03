const express = require("express");
const fs = require("fs")
const nodemailer = require("nodemailer");
const { MongoClient } = require("mongodb");
const {onRequest} = require("firebase-functions/v2/https");
const {initializeApp} = require("firebase-admin/app");
require("dotenv").config();

initializeApp();

const app = express();
const port = 8100;
const godaddyPassword = process.env.GODADDY_PASSWORD;
const godaddyEmail = process.env.GODADDY_EMAIL;
const godaddySMTP = process.env.GODADDY_SMTP;
const mongoURI = process.env.MONGO_URI;
const ytDemoLink = process.env.YT_VIDEO_URL;
const client = new MongoClient(mongoURI);
var mTemplate;

fs.readFile("./content/mail.html", "utf-8",(err, mailTemplate) => {
    if(err){
        console.error("Cannot read html mail template")
        return
    }
    mTemplate = mailTemplate
})

app.use(express.json());

app.post("/share-demo", (req, res) => {
    const { firstname,lastname ,email, company } = req.body;
    var shareDemoMailTemplate = mTemplate;
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
    subject: "DevXStudio Demo",
    html: shareDemoMailTemplate
        .replace(/{{firstname}}/g, firstname)
        .replace(/{{ytDemoLink}}/g, ytDemoLink),
    attachments: [
        {
            filename: "questions.pdf",
            path: "./content/godaddy-studiodevx.pdf",
        }
    ]
  };

    transporter.sendMail(mailOptions, (error, info) => {
        if (error) {
            console.log(error);
            res.status(500).send("Something went wrong.");
        } else {
            console.log("Email sent: " + info.response);
            res.status(200).send("Email successfully sent to recipient!");
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
        res.status(500).send("Something went wrong.");
    }
});

app.post("/waitlist", (req, res) => {
    const { firstname, email } = req.body;
    try {
        client.db("devx-contact").collection("waitlist").insertOne({
            firstname: firstname,
            email: email
        })
    } catch (error) {
        res.status(500).send("Something went wrong.");
    }
    res.status(200).send("Contact waitlist successfully added!");
});

exports.waitlist = onRequest({cors: true},app);