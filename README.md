# Deploy Your Code On a Docker Container Using Jenkins on AWS

![diagram-export-4-22-2025-7_09_24-PM](https://github.com/user-attachments/assets/44669afe-7ed3-46cd-b890-be4ec4b88dde)

Description 

## Agenda
1. Setup jenkins
2. Setup and configure git
3. Integration github with jenkins
4. Setup docker host
5. Integrate docker with jenkins
6. Automate the build and deploy process using jenkins
7. test the deployment

## Prequisites
1. Amazon Web Services (AWS)
2. SonarQube
3. Github account with the source code

## STEP 1 : Setup Jenkins Server on AWS EC2
1. Setup the linux instance
2. Install Java
3. Install Jenkins
4. Configure and Start Jenkins
5. Access web UI on PORT 8080

Login ke halaman Amazon Management Console, buka EC2 Dashboard, dan klik Launch Instance:
![Screenshot from 2025-04-22 19-20-46](https://github.com/user-attachments/assets/8702d2a9-4229-422e-accf-04ddc4cb8414)

Setelah meng-klik launch instance, berikan nama untuk server jenkins:
![Screenshot from 2025-04-22 19-31-18](https://github.com/user-attachments/assets/ad91c197-bcc9-4c93-a925-c1d1753be85a)

Selanjutnya memilih sistem operasi yang akan digunakan untuk server jenkins, disini saya memakai OS Ubuntu 22.04:
![Screenshot from 2025-04-22 19-39-41](https://github.com/user-attachments/assets/128d55a2-7fb7-41be-9441-144bade883d9)

Pilih instance type yang akan digunakan pada server jenkins:
![image](https://github.com/user-attachments/assets/07b7ff07-83b2-4e44-b21e-b77134ef4e09)

Selanjutnya membuat key pairs untuk mengakses server jenkins di terminal, buat key-pair baru:
![Screenshot from 2025-04-22 19-43-46](https://github.com/user-attachments/assets/cb999d63-c730-453f-bf8e-17c5a06e30ad)

Berikan nama untuk key-pair lalu pilih key pair 'RSA' dan file format nya '.pem'
![Screenshot from 2025-04-22 19-44-51](https://github.com/user-attachments/assets/c4418687-1df5-4c41-b23b-c5aeb78358b2)

Selanjutnya membuat storage, setelah itu klik launch instance:
![image](https://github.com/user-attachments/assets/dd64f31c-a3f0-43e7-97b7-1e8821fe7da3)

Buka terminal pada linux lalu ketikan perintah berikut ini untuk remote server aws ec2
```
ssh -i [file-ssh] [nama-user-server@ipaddress-server]
```
![Screenshot from 2025-05-02 16-16-44](https://github.com/user-attachments/assets/e1209936-06a1-4f20-ae08-a78f87d2eaf8)

Ubah nama server menjadi jenkins dengan cara mengetikkan perintah
```
sudo hostnamectl set-hostname jenkins
/bin/bash
```
![Screenshot from 2025-05-02 16-21-47](https://github.com/user-attachments/assets/b8327b33-dcdd-45f4-9166-068c566d0330)

Setelah merubaha nama server selanjutnya adalah menginstall java pada server jenkins, ketikkan perintah berikut ini untuk menginstall java di terminal
```
sudo apt install openjdk-17-jdk openjdk-17-jre

```
![Screenshot from 2025-04-30 14-40-26](https://github.com/user-attachments/assets/bc90eb53-4fd7-4248-9251-67f8798ad094)

Setelah berhasil install, ketikkan perintah berikut ini untuk memeriksa apakah penginstalan java sudah berhasil
![Screenshot from 2025-05-02 16-30-36](https://github.com/user-attachments/assets/362f2534-dbbc-4df4-8ced-106835d76209)

Setelah JAVA berhasil di install di server jenkins, selanjutnya adalah install jenkins di server. Kunjungsi website resmi jenkins lalu cari dokumentasi penginstalan jenkins pada OS Linux Ubuntu 
![Screenshot from 2025-05-02 16-35-30](https://github.com/user-attachments/assets/d368b169-3514-4f52-af8a-495c494d77e4)

Pilih perintah Long Term Support release, copy dan paste ke terminal server 
![Screenshot from 2025-04-30 14-48-54](https://github.com/user-attachments/assets/eee9930b-ed3c-4e2b-b719-8f1d270fd562)

Setelah proses instalasi jenkins selesai, selanjut nya aktifkan jenkins dan lihat status jenkins dengan perintah berikut ini 
```
sudo systemctl start jenkins
sudo systemctl status jenkins
```
![Screenshot from 2025-05-02 16-41-09](https://github.com/user-attachments/assets/8c360c42-93fe-4ff9-971c-3fb40e240143)

Sekarang akses jenkins dengan memasukkan ip server jengkins ditambah dengan port default nya yaitu 8080
```
52.77.231.48:8080
```
![Screenshot from 2025-04-30 14-51-23](https://github.com/user-attachments/assets/39b490a8-0359-43c7-8d29-f4fc59b67492)

Customize halaman jenkins, pilih install yang di rekomendasikan

![customize jenkins](https://github.com/user-attachments/assets/8bbf21d9-eb0c-4b95-b13b-d6b917ac1715)

Buat buat akun admin user dan isi kolom yang ada pada gambar berikut

![WhatsApp Image 2025-05-02 at 17 36 57](https://github.com/user-attachments/assets/32723058-f463-4deb-86e7-cd9de7412013)

Setelah mengisi data pada proses sebelumnya, jenkins siap untuk digunakan

![ready](https://github.com/user-attachments/assets/4929b704-33b0-4e35-b5d3-b63f7c27d610)

## STEP 2 : Integrate Jenkins dengan Github

Buka halaman jenkins menggunakan alamat IP pada server jenkins + PORT 8080, lalu pilih New Item untuk membuat project baru

![Screenshot from 2025-05-05 15-33-02](https://github.com/user-attachments/assets/8e012a2a-9623-4dcd-8e53-9375b866abf9)

Buat nama project dan type yang kita pilih disini adalah freestyle project lalu klik OK

![Screenshot from 2025-05-05 15-43-42](https://github.com/user-attachments/assets/936f48d4-2578-4e14-ae01-34a537795b91)

Setelah itu klik Source Code Management lalu pilih GIT

![Screenshot from 2025-05-05 15-53-54](https://github.com/user-attachments/assets/5d71bd72-4199-4298-affc-65dcdb12eb77)

Coppy URL repository github 

![Screenshot from 2025-05-05 15-56-18](https://github.com/user-attachments/assets/14d0f743-a095-45a7-84ae-1541d5f90954)

Paste URL yang sudah di salin tadi ke kolom berikut 

![Screenshot from 2025-05-05 16-03-45](https://github.com/user-attachments/assets/7fbfb28d-fb83-4fc4-a5c9-c0dc33f28aa3)

Lalu sesuaikan branches to build dengan branch kerja kalian 

![Screenshot from 2025-05-05 16-05-20](https://github.com/user-attachments/assets/483d175b-5002-4633-ae33-ba0d044322a8)

Selanjutnya ceklis pada bagian Github hook trigger fot GITSM Polling agar setiap perubahan yang terjadi pada repository Github, jenkins akan segera menangkap perubahan tersebut lalu Save

![image](https://github.com/user-attachments/assets/235cbf4b-06b1-4143-a82c-11f1ae8ad91f)

Selanjutnya adalah membuat webhook di github. untuk membuat webhook github, kita harus pergi ke pengaturan pada repository kita lalu pilih Webhook lalu pilih Add Webhook

![Screenshot from 2025-05-05 16-20-02](https://github.com/user-attachments/assets/55a28926-6b00-472f-a0a7-ce591822220c)

Setelah itu isi kolom Payload URL dengan alamat IP Jenkins + PORT yang terhubung lalu tambahkan endpoint /github-webhook agar GitHub tahu ke mana harus mengirim notifikasi saat terjadi event (seperti push, pull request, dll).

![Screenshot from 2025-05-05 16-24-48](https://github.com/user-attachments/assets/659368ca-d426-4781-8ceb-20f295b78bfb)

Scroll kebawah dan centang bagian Let me select individual events 

![Screenshot from 2025-05-05 16-35-32](https://github.com/user-attachments/assets/f86cb77b-05ce-4eb9-aa06-96618fff271a)

lalu centang bagian pull request dan add webhook

![image](https://github.com/user-attachments/assets/09636357-d9cc-41dc-9935-fe1f9f52e2d7)

## STEP 3 : Integrate Jenkins dengan SonarQube

Buat server baru di AWS EC2 dengan mengikuti langkah-langkah diatas, setelah berhasil membuat server baru install openjdk java

![Screenshot from 2025-05-23 14-36-59](https://github.com/user-attachments/assets/0d71714c-2686-4f5b-829c-ad5c70713ed8)

Setelah berhasil menginstall openjdk java, selanjutnya adalah download SonarQube 

![Screenshot from 2025-05-23 14-32-54](https://github.com/user-attachments/assets/8dc29439-e425-488d-a0f5-6d2379c50db7)

Setelah berhasil download SonarQube, selanjutnya adalah menginstall UNZIP untuk mengextract file SonarQube yang sudah di download tadi

![Screenshot from 2025-05-23 14-34-01](https://github.com/user-attachments/assets/69e5215c-1b35-49d2-844c-5ea5dbbbeab3)

Selanjutnya adalah UNZIP file SonarQube yang sudah di download tadi 

![Screenshot from 2025-05-23 14-34-35](https://github.com/user-attachments/assets/9ca9be46-3048-47fb-93b2-86c59dddca17)

Setelah di unzip, selanjutnya buka folder SonarQube dan jalankan sonarqube dengan perintah ./sonar.sh start

![Screenshot from 2025-05-23 14-35-24](https://github.com/user-attachments/assets/d610a4e3-01bd-459a-942f-4dd825842cbf)

Selanjutnya adalah mengizinkan port default dari Sonarqube PORT 9000 di Inbound Rules AWS EC2 

![Screenshot from 2025-05-23 14-39-24](https://github.com/user-attachments/assets/1d6067c9-3596-4c92-a18d-765fa7b9fc4b)

Setelah mengizinkan PORT 9000 di Inbound Rules server SonarQube, selanjutnya akses halaman SonarQube di browser dengan IP-SERVER + PORT:9000.

![Screenshot from 2025-05-23 15-44-19](https://github.com/user-attachments/assets/f60fcd48-19e3-4549-bc26-b49cb52e3492)

Setelah berhasil mengkonfigurasi user sonarqube selanjutnya adalah membuat project SonarQube, pada project ini menggunakan "Create a local project"

![Screenshot from 2025-05-23 15-45-27](https://github.com/user-attachments/assets/8eb490c4-caba-437b-8d83-8fd529e8803d)

Lalu masukkan nama project dan branch name dari repository 

![Screenshot from 2025-05-23 15-45-38](https://github.com/user-attachments/assets/6f51ac00-b139-41ae-870e-f05311a70dd2)

Selanjutnya pilih global setting 

![Screenshot from 2025-05-23 15-45-48](https://github.com/user-attachments/assets/fd6d9933-18bb-4e6f-999c-1f33f115a06b)

Langkah selanjutnya adalah pilih "with jenkins" untuk analysis method 

![Screenshot from 2025-05-23 15-46-02](https://github.com/user-attachments/assets/0da2a667-17ea-4746-831c-62b58b33cf73)

Lalu pilih Github untuk DevOps Platform

![Screenshot from 2025-05-23 15-46-12](https://github.com/user-attachments/assets/9eab01de-5b54-4bff-8318-56656ebc2820)

Install plugin SonarQube di Jenkins

![Screenshot from 2025-05-23 15-50-06](https://github.com/user-attachments/assets/f75a3b33-1abf-4f12-b23c-7186c5ee03cf)

Setelah berhasil menginstall plugin SonarQube, selanjutnya Add SonarQube installations di pengaturan Manage Jenkins > System

![Screenshot from 2025-05-23 15-51-00](https://github.com/user-attachments/assets/c91e9a05-30d5-4646-bc56-e07123804277)

Lalu konfigurasi Nama Server, Server URL diisi dengan alamat IP dan PORT yang digunakan SonarQube, dan untuk Server Authentication Token pilih ADD > Jenkins

![Screenshot from 2025-05-23 15-51-28](https://github.com/user-attachments/assets/6d0a6964-588f-4f97-8869-2888d1282a7c)

Untuk menambahkan Server Authentication Token, pergi ke My Account di pojok kanan atas, lalu di menu security isi nama token, untuk type nya Global Analysis token dan expired lalu klik generate 

![Screenshot from 2025-05-23 15-52-15](https://github.com/user-attachments/assets/eb4ae3fe-efc6-453e-b809-299b498b2a3a)

Copy token nya lalu pergi ke menu Server Authentication Server di Jenkins, di Jenkins Credentials Provider Jenkins pilih Secret Text dan masukkan token tadi lalu Save pengaturan

![Screenshot from 2025-05-23 15-53-17](https://github.com/user-attachments/assets/f810bd93-38fd-416b-8b2e-5a4fdd3ace5a)

Scroll kebawah pada Analyze your project with Jenkins di SonarQube, lalu pada menu Create a Jenkinsfile pilih Other(for Go, Python, PHP), maka akan muncul code sonar-project-properties. copy code tersebut 

![Screenshot from 2025-05-23 15-54-37](https://github.com/user-attachments/assets/98330b34-6787-46fe-a8cb-adf2432acf60)

Selanjutnya pergi ke project kita, lalu pilih Environment dan tambahkan build step lalu pilih Execute SonarQube Server, lalu paste code sonar-project-properties tadi di analysis properties lalu save

![Screenshot from 2025-05-23 21-11-48](https://github.com/user-attachments/assets/f1d64a15-f15e-40cf-b53d-61e0f6dff5fb)

Selanjutnya build ulang project dan lihat hasil scan oleh SonarQube di dashboard dengan mengakses <ip-server-sonarqube>+<port> di web browser. erikut adalah hasil dari scan source code

![image](https://github.com/user-attachments/assets/691a41c0-c69a-4c70-9198-9f3782ebe372)

## STEP 4 : Integrate Jenkins dengan Docker

Akses server docker menggunakan SSH dan gunakan perintah "sudo apt get update" untuk update pembaharuan yang tersedia di server 

![Screenshot from 2025-06-01 19-11-57](https://github.com/user-attachments/assets/945a6f74-320c-4b36-8376-154ab7e1d5ea)

Setelah berhasil di update, selanjutnya download docker di web berikut ini :https://docs.docker.com/engine/install/ubuntu/ .pastikan untuk os yang digunakan di server sesuai dengan instalasi docker 

![Screenshot from 2025-06-01 19-12-16](https://github.com/user-attachments/assets/14233d9f-a0c3-446e-af96-96ba7fd27c2d)

Copy perintah 1. Setup Docker apt repository dan paste di terminal 

![Screenshot from 2025-06-01 19-12-38](https://github.com/user-attachments/assets/ceac9a03-ce53-4c23-9540-7d0e4264cb5b)

Selanjutnya install the docker packages di terminal dan verifikasi instalasi docker dengan menggunakan perintah "sudo docker run hello-world"

![Screenshot from 2025-06-01 19-13-16](https://github.com/user-attachments/assets/f87836ff-965e-4f48-8b45-41e78b102481)

Selanjutnya adalah mengkonfigurasi file SSHD_CONFIG agar server Docker dapat diakses menggunakan password dengan cara masuk sebagai root dan ketikkan perintah "nano /etc/ssh/sshd_config"

Konfigurasi PubkeyAuthentication apabila sebelumnya bernilai "NO" ubah menjadi "YES"

![Screenshot from 2025-06-01 19-16-56](https://github.com/user-attachments/assets/4d863ffc-d203-4613-943e-c683bc66c1cb)

Selanjutnya pada bagian PasswordAuthentication ubah dari "NO" menjadi "YES"

![Screenshot from 2025-06-01 19-17-09](https://github.com/user-attachments/assets/33167b22-f23d-49af-a60e-54f03b75ec62)

Setelah berhasil di konfigurasi, selanjutnya restart sshd menggunakan perintah "systemctl restart sshd"

![Screenshot from 2025-06-01 19-17-32](https://github.com/user-attachments/assets/518b540f-b293-423c-926b-dde94606b3ef)

Selanjutnya adalah membuat password untuk user root/ubuntu dengan perintah passwd 

![Screenshot from 2025-06-01 19-19-54](https://github.com/user-attachments/assets/a60fa377-b527-4f32-a1bf-cc66bd2aa023)

Buka server jenkins dan akses sebagai root dan buat ssh-keygen agar jenkins dapat mengakses server docker

![Screenshot from 2025-06-01 19-21-15](https://github.com/user-attachments/assets/91c594b2-73c1-40ac-804a-2363abc6f2c2)

Salin isi dari id_rsa.pub mulai dari ssh-rsa sampai user@user

![image](https://github.com/user-attachments/assets/0b1a0c41-8e6b-4ae8-af30-a2b457b3e2f6)

dan simpan id_rsa.pub ke server docker dan di simpan di folder /home/ubuntu/.ssh/authorized keys

![image](https://github.com/user-attachments/assets/eaafc15b-8e24-4b3c-9b48-6e775f18bdae)

server docker berhasil di akses menggunakan password

![image](https://github.com/user-attachments/assets/49d0977f-7a24-4a89-bd7a-1c52d8168d7d)

Selanjutnya install docker-compose di server docker menggunakan perintah berikut ini

![Screenshot from 2025-06-01 23-03-10](https://github.com/user-attachments/assets/b357da9c-cc1d-48a5-8a99-6174114d248d)

Setelah berhasil diinstall, selanjutnya adalah mengkonfigurasi Jenkins. Buka server jenkins dan masuk ke menu Dashboard/Manage Jenkins/System. tambahkan pada menu server group list dan masukkan group name dengan nama Docker, SSH PORT default ssh yaitu 22, username dari server docker = ubuntu dan password yang sudah dibuat

![image](https://github.com/user-attachments/assets/c36c6426-97cf-4fa6-bbb6-f421905f5dc2)

Lalu pada bagian server list masukkan nama server-name dan server ip. server IP memakai IP server DOCKER selanjutnya save

![image](https://github.com/user-attachments/assets/c2d9ea07-3c7e-4ce2-8c16-2339fe4bde12)






























