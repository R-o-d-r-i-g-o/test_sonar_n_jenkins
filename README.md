# test_sonar_n_jenkins

## to use swagger:
<pre>
  $ swag fmt  // format comments
  $ swag init // generate swagger routes
</pre>

## to use sonar:
<pre>
  $ docker run -d --name sonarqube -e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true -p 9000:9000 sonarqube:latest
</pre>
